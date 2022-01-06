package main

import (
	"classic"
	"fmt"
	"hangman_web/structure"
	"hangman_web/utils"
	"html/template"
	"net/http"
)

var stock1 = classic.GetRandomWord()

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.gohtml"))
	err := tmpl.Execute(w, manager())
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

//func Hangman(w http.ResponseWriter, r *http.Request) {
//	tmpl := template.Must(template.ParseFiles("index.gohtml"))
//	}

var stock2 = classic.GetRandomWord()

func manager() structure.Stock {
	target := stock2
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

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/")
	http.ListenAndServe(":8000", server)
}
