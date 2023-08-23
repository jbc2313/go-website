package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type PageDetail struct {
	Title string
	Body []string
	Footer string
}

func greet(w http.ResponseWriter, r *http.Request) {
	data := PageDetail {
		Title: "GoSite",
		Body: []string{"This site is running on Go","It is generally faster then Node.js","Go has concurrency, allowing multiple tasks at once to take place."},
		Footer: "Golang",
	}
	tmpl.Execute(w, data)
}

func about(w http.ResponseWriter, r *http.Request) {
	data := PageDetail {
		Title: "About",
		Body: []string{"This site is built with only the go std lib"},
		Footer: "Golang",
	}
	tmpl.Execute(w, data)
};

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml", "templates/header.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", greet)
	mux.HandleFunc("/about", about)

	log.Fatal(http.ListenAndServe(":8090", mux))

}
