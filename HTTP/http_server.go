package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func dobro(n int) int {
	return n * 2
}

type user struct {
	Name  string
	Email string
}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u1 := user{"Maria Joaquina", "joaoavilaperasol@hotmail.com"}

		templates.ExecuteTemplate(w, "index.html", u1)
	})

	http.HandleFunc("/dobro", func(w http.ResponseWriter, r *http.Request) {
		var param int = 4
		var number int = dobro(param)
		var text string = fmt.Sprintf("O dobro de %d é %d", param, number)

		w.Write([]byte(text))
	})

	log.Fatal(http.ListenAndServe(":5555", nil))
}
