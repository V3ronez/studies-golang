package main

import (
	"log"
	"net/http"
	"text/template"
)

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	u := usuario{"Henrique", "henrique@gmail.com"}
	// u := usuario{Nome: "Henrique", Email: "henrique@gmail.com"}
	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
