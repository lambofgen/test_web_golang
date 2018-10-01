package main

import (
	"net/http"

	"github.com/lambofgen/testweb/route"
)

func main() {
	r := route.Create()
	http.ListenAndServe(":8080", r)
}
