package dbService

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_project/pavilions"
	"go_project/shelters"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	_ "github.com/lib/pq"
)

var (
	lock = new(sync.RWMutex)
	DB   *sql.DB
)

func InsertSelter(id uint64, address string) error {
	_, err := DB.Exec("INSERT INTO shelter VALUES ($1, $2)", id, address)
	if err != nil {
		return err
	}
	return nil
}

func InsertPavilions(id uint64, name string, shelterID uint64, count uint64) error {
	_, err := DB.Exec("INSERT INTO pavilions VALUES ($1, $2, $3,$4)",
		id, name, shelterID, count)
	if err != nil {
		return err
	}
	return nil
}

func SelectAllPavilions() ([]pavilions.Pavilion, error) {
	rows, err := DB.Query("SELECT id,shelter_id,count,name FROM pavilions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]pavilions.Pavilion, 0)
	var pavilion pavilions.Pavilion
	for rows.Next() {
		err = rows.Scan(&pavilion.Number, &pavilion.ID, &pavilion.Count, &pavilion.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, pavilion)
	}
	return result, nil
}

func SelectPavilion(id uint64) (*pavilions.Pavilion, error) {
	rows, err := DB.Query("SELECT number,name,id,count FROM pavilions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result pavilions.Pavilion
	err = rows.Scan(&result.Number, &result.Name, &result.ID, &result.Count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func SelectAllShelters() ([]shelters.Shelter, error) {
	rows, err := DB.Query("SELECT id,address FROM shelter")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]shelters.Shelter, 0)
	var shelter shelters.Shelter
	for rows.Next() {
		err = rows.Scan(&shelter.ID, &shelter.Address)
		if err != nil {
			return nil, err
		}
		result = append(result, shelter)
	}
	return result, nil
}

func SelectShelter(id uint64) (*shelters.Shelter, error) {
	rows, err := DB.Query("SELECT id,address FROM shelter WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result shelters.Shelter
	err = rows.Scan(&result.ID, &result.Address)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func DeleteU(name string, id uint64) error {
	var err error
	switch name {
	case "shelter":
		_, err = DB.Exec("DELETE FROM shelter WHERE id = $1", id)
	case "pavilions":
		_, err = DB.Exec("DELETE FROM pavilions WHERE id = $1", id)
	}

	if err != nil {
		return err
	}
	return nil
}

func getID(args *url.Values) (uint64, error) {
	var err error

	if id_str := args.Get("id"); id_str != "" {
		var id uint64
		if id, err = strconv.ParseUint(id_str, 10, 64); err != nil {
			return 0, err
		}
		return uint64(id), nil
	}
	return 0, fmt.Errorf("Не указан идентификатор")
}

func getShelterId(args *url.Values) (uint64, error) {
	var err error

	if id_str := args.Get("shelterId"); id_str != "" {
		var id uint64
		if id, err = strconv.ParseUint(id_str, 10, 64); err != nil {
			return 0, err
		}
		return uint64(id), nil
	}
	return 0, fmt.Errorf("Не указан id питомника")
}

func getCount(args *url.Values) (uint64, error) {
	var err error

	if id_str := args.Get("count"); id_str != "" {
		var id uint64
		if id, err = strconv.ParseUint(id_str, 10, 64); err != nil {
			return 0, err
		}
		return uint64(id), nil
	}
	return 0, fmt.Errorf("Не указано количество животных")
}

func getTable(args *url.Values) (string, error) {

	if table := args.Get("table"); table != "" {
		return table, nil
	}
	return "", fmt.Errorf("Не указана таблица")
}

func getName(args *url.Values) (string, error) {

	if table := args.Get("name"); table != "" {
		return table, nil
	}
	return "", fmt.Errorf("Не указано животное")
}

func getAddress(args *url.Values) (string, error) {

	if table := args.Get("address"); table != "" {
		return table, nil
	}
	return "", fmt.Errorf("Не указан адрес")
}

func Delete(req *http.Request) ([]byte, error) {
	args := req.URL.Query()

	id, err := getID(&args)
	if err != nil {
		return nil, err
	}

	name, err := getTable(&args)
	if err != nil {
		return nil, err
	}

	lock.Lock()
	defer lock.Unlock()
	DeleteU(name, id)
	return json.Marshal(map[string]bool{"OK": true})
}

func CreateU(tableName string, id uint64) error {
	var err error
	switch tableName {
	case "shelter":
		_, err = DB.Exec("DELETE FROM shelter WHERE id = $1", id)
	case "pavilions":
		_, err = DB.Exec("DELETE FROM pavilions WHERE id = $1", id)
	}

	if err != nil {
		return err
	}
	return nil
}

func Create(req *http.Request) ([]byte, error) {
	args := req.URL.Query()

	tableName, err := getTable(&args)
	if err != nil {
		return nil, err
	}

	id, err := getID(&args)
	if err != nil {
		return nil, err
	}

	switch tableName {
	case "shelter":
		address, _ := getAddress(&args)
		err = InsertSelter(id, address)
	case "pavilions":
		name, _ := getName(&args)
		shID, _ := getShelterId(&args)
		count, _ := getCount(&args)
		err = InsertPavilions(id, name, shID, count)
	}

	lock.Lock()
	defer lock.Unlock()

	return json.Marshal(map[string]bool{"OK": true})
}

func Select(req *http.Request) ([]byte, error) {
	args := req.URL.Query()

	name, err := getTable(&args)
	if err != nil {
		return nil, err
	}

	id, err := getID(&args)
	if err != nil {
		return nil, err
	}

	lock.Lock()
	if name == "shelter" {
		//var r, _ = SelectShelter(id)
		var r, _ = SelectAllShelters()
		if r == nil {
			return json.Marshal(map[string]string{"result": "it is nil m*ka"})
		}
		//fmt.Print((*r).String())
		return json.Marshal(r)
		//return json.Marshal(map[string]string{"result": "rghfjkhkjffh"})
	} else if name == "pavilions" {
		var r, _ = SelectPavilion(id)
		fmt.Print((*r).String())
		return json.Marshal((*r).String())
	} else {
		return nil, nil
	}
	defer lock.Unlock()

	return json.Marshal(map[string]bool{"Bad request": true})
}

func Live(rw http.ResponseWriter, req *http.Request) {
	var err error
	var res []byte

	switch req.Method {
	case "POST", "PUT":
		res, err = Create(req)
		//	case "PATCH":
		//		res, err = Update(req)
	case "DELETE":
		res, err = Delete(req)
	default:
		res, err = Select(req)
	}

	if err != nil {
		data, _ := json.Marshal(map[string]interface{}{"OK": false, "Error": err.Error()})
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(data)
		fmt.Printf("[ %s 400 ] %s : %s", req.Method, req.URL, err)
		return
	}

	if res == nil && len(res) == 0 {
		http.NotFound(rw, req)
		fmt.Printf("[ %s 404 ] %s", req.Method, req.URL)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(res)
	fmt.Printf("%s %s %d", req.Method, req.URL, len(res))
}
