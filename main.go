package main

import (
	"net/http"

	"./web"
)

func main() {
	http.ListenAndServe(":3000", web.NewRouter())
}
