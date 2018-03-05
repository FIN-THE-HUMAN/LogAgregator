package main

/*
import (
	"fmt"
	"go_project/DB"
	"go_project/dbService"
	"go_project/pavilions"
	"go_project/shelters"
	"go_project/stream"
	"net/http"
)

func main() {
	var err error
	dbService.DB, err = DB.Connect("localhost", "5432", "postgres", "22821810", "laba", "disable")
	defer dbService.DB.Close()

	var pavilions []pavilions.Pavilion
	pavilions, err = dbService.SelectAllPavilions()
	fmt.Println(pavilions)

	fmt.Println()

	var shelters []shelters.Shelter
	shelters, err = dbService.SelectAllShelters()
	fmt.Println(shelters)

	//	v, _ := dbService.SelectShelter(1)
	//	print((*v).String())
	//err = dbService.DeleteU("shelter", 5)

	//	http.HandleFunc("/insertShelter", dbService.InsertSelter)
	//	http.HandleFunc("/insertPavilion", InsertPavilions)
	//	http.HandleFunc("/deleteShelter", Index)
	http.HandleFunc("/delete", dbService.Live)
	http.HandleFunc("/select", dbService.Live)
	http.HandleFunc("/create", dbService.Live)
	http.HandleFunc("/insert", dbService.Live)

	//	http.HandleFunc("/selectShelter", Index)
	//	http.HandleFunc("/selectPavilion", Index)

	http.ListenAndServe(":3000", nil)
	stream.Catch(err)

}
*/

import (
	"fmt"
	"go_project/error"
)

func main() {
	var err error.ValidateError
	err.AddMessage("Old Error")
	err.AddMessage("New Error")
	err.AddMessage("Last Error")
	fmt.Print(err.Message)
}
