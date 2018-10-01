package route

import (
	"net/http"

	"github.com/lambofgen/testweb/controllers"
)

func Create() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.HomeController)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return mux
}
