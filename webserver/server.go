//package webserver
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	homeTmpl = template.Must(template.ParseFiles("index.html.tmpl"))
)

func homepage(response http.ResponseWriter, request *http.Request) {
	err := homeTmpl.Execute(response, map[string]string{
		"Name": request.URL.Query().Get("name"),
	})

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homepage)
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		log.Fatal(err)
	}
}
