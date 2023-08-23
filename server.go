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
		Body: []string{"This site is running on Go","It is 10x faster then Node.js"},
		Footer: "Golang",
	}
	tmpl.Execute(w, data)
}


func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", greet)

	log.Fatal(http.ListenAndServe(":8090", mux))

}
