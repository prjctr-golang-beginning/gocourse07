package main

import (
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Title string
	Body  string
}

var templates = template.Must(template.ParseFiles(
	"boilerplate/templates/header.html",
	"boilerplate/templates/nav.html",
	"boilerplate/templates/index.html",
	"boilerplate/templates/footer.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string, page *Data) {
	err := templates.ExecuteTemplate(w, tmpl, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	page := &Data{Title: "Home page", Body: "Welcome to our brand new home page."}
	renderTemplate(w, "index", page)
}

func AboutHandler(w http.ResponseWriter, _ *http.Request) {
	page := &Data{Title: "About page", Body: "This is our brand new about page."}
	renderTemplate(w, "index", page)
}

func main() {
	log.Println(`--- Handlers defined`)
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", AboutHandler)
	log.Println(`--- Webserver started`)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
