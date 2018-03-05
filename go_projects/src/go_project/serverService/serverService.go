package serverService

import (
	"net/http"

	_ "github.com/lib/pq"
)

func viewHandlerInsert(w http.ResponseWriter, r *http.Request) {

}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
