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
		Data  string
	}{Title: "ok", Data: classic.HideLetters(classic.GetRandomWord())},
	)

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	letter := r.FormValue("letter")
	fmt.Fprintf(w, "Letter = %s\n", letter)
}

func main() {

	server := http.NewServeMux()
	// url http://localhost:8000/
	server.HandleFunc("/", Home)
	server.HandleFunc("/hangman", Hangman)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/hangman")
	http.ListenAndServe(":8000", server)
}
