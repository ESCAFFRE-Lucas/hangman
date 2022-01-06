package main

import (
	"classic"
	"fmt"
	"hangman_web/structure"
	"hangman_web/utils"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.gohtml"))
	_ = tmpl.Execute(w, manager())
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	target := classic.GetRandomWord()
	tmpl := template.Must(template.ParseFiles("index.gohtml"))
	err := tmpl.Execute(w, structure.Stock{
		Title: "Hangman", Right: []string{}, Wrong: []string{}, Attempts: 10, TargetWord: target, CurrentWord: classic.InitWord(target)},
	)
	if err != nil {
		fmt.Println(err)
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	letter := r.FormValue("letter")
	fmt.Fprintf(w, "letter :%s\n", letter)
}

func manager() structure.Stock {
	target := classic.GetRandomWord()
	data := utils.LoadFile()
	fmt.Println(data)
	if data.TargetWord == "" {
		data = structure.Stock{
			Title:       "Hangman",
			Right:       []string{},
			Wrong:       []string{},
			Attempts:    10,
			TargetWord:  target,
			CurrentWord: classic.InitWord(target),
		}
		utils.SaveInFile(data)
	}
	return data
}

func main() {
	server := http.NewServeMux()
	// url http://localhost:8000/
	server.HandleFunc("/", Home)
	server.HandleFunc("/hangman", Hangman)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/")
	http.ListenAndServe(":8000", server)
}
