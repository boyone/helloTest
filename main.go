package main

import (
	"net/http"

	"github.com/boyone/helloTest/web"
)

func main() {
	http.ListenAndServe(":3000", web.NewRouter())
}
