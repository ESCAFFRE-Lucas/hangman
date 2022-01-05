package main

import (
	"classic"
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	_ = tmpl.Execute(w, nil)
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.gohtml"))
	_ = tmpl.Execute(w, struct {
		Title string
		Data  []string
	}{Title: "ok", Data: []string{classic.GetRandomWord(), classic.Input(), " ", "p", "d"}},
	)
}

func main() {

	server := http.NewServeMux()
	// url http://localhost:8000/
	server.HandleFunc("/", Home)
	server.HandleFunc("/gohtml", Hangman)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/gohtml")
	http.ListenAndServe(":8000", server)
}
