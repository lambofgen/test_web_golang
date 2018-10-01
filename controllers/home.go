package controllers

import (
	"html/template"
	"net/http"

	"github.com/lambofgen/testweb/services"
	"github.com/lambofgen/testweb/views"
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
