package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", func(res http.ResponseWriter, req *http.Request) {
		v := mux.Vars(req)
		fmt.Fprintf(res, "Hello, %s", v["name"])
	}).Methods("GET")
	return r
}
