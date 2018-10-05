package main

import (
	"net/http"

	"github.com/lambofgen/test_web_golang/route"
)

func main() {
	r := route.Create()
	http.ListenAndServe(":8080", r)
}
