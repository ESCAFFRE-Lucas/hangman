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
	data, _ := manager(nil)
	_ = tmpl.Execute(w, data)
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	letter := r.FormValue("letter")
	manager(&letter)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	if classic.IsNotALetter(letter) {
		http.Redirect(w, r, "/errors", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "../", http.StatusSeeOther)
	}
}

var AttemptLeft = 10

func manager(input *string) (structure.Stock, bool) {
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
		if *input == data.TargetWord {
			AttemptLeft = 10
			utils.SaveInFile(structure.Stock{})
		} else {
			data.Attempts = AttemptLeft - len(data.Wrong)
			utils.SaveInFile(data)
		}
	}
	if data.Attempts <= 0 {
		AttemptLeft = 10
		utils.SaveInFile(structure.Stock{})
	} else if data.CurrentWord == data.TargetWord {
		AttemptLeft = 10
		utils.SaveInFile(structure.Stock{})
		return data, true
	}
	return data, false
}

func DisplayErrors(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Errors/errors-page.gohtml"))
	_ = tmpl.Execute(w, nil)
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("start/startgame.gohtml"))
	_ = tmpl.Execute(w, nil)
}

func ScoreBoard(r *http.Request, count int) map[string]int {
	username := r.FormValue("username")
	score := map[string]int{username: count}
	return score
}

func main() {
	server := http.NewServeMux()
	// url http://localhost:8000/
	server.HandleFunc("/", Home)
	server.HandleFunc("/hangman", Hangman)
	server.HandleFunc("/errors", DisplayErrors)
	server.HandleFunc("/start", StartGame)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/start")

	http.ListenAndServe(":8000", server)
}
