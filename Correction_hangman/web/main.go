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
	_ = tmpl.Execute(w, manager(nil))
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	letter := r.FormValue("letter")
	manager(&letter)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	http.Redirect(w, r, "../", http.StatusSeeOther)
}

var AttemptLeft = 10

func manager(input *string) structure.Stock {
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
	if input != nil {
		classic.HandleInput(data.TargetWord, *input, &data.CurrentWord, &data.Right, &data.Wrong, data.Attempts)
		if len(*input) > 1 {
			AttemptLeft -= 1
		}
		data.Attempts = AttemptLeft - len(data.Wrong)
		utils.SaveInFile(data)
	}
	if data.Attempts <= 0 || data.CurrentWord == data.TargetWord {
		AttemptLeft = 10
		utils.SaveInFile(structure.Stock{})
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
