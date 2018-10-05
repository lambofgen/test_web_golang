package controllers

import (
	"html/template"
	"net/http"

	"github.com/lambofgen/test_web_golang/services"
	"github.com/lambofgen/test_web_golang/views"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	data, err := services.GetResponse()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	t, err := template.ParseFiles(views.HOME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, data)
}
