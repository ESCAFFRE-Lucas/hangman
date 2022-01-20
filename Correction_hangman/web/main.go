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
	data, _ := manager(nil, nil)
	_ = tmpl.Execute(w, data)

}

var AttemptLeft = 10

func Hangman(w http.ResponseWriter, r *http.Request) {
	letter := r.FormValue("letter")
	_, gameWon := manager(&letter, nil)
	ok := true
	if gameWon == &ok {
		http.Redirect(w, r, "/endwin", http.StatusSeeOther)
	} else if gameWon != &ok {
		http.Redirect(w, r, "/endlose", http.StatusSeeOther)
	}
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

func Redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func manager(input *string, difficulty *string) (structure.Stock, *bool) {
	target := classic.GetRandomWord(difficulty)
	data := utils.LoadFile()
	vrai := true
	faux := false
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
		return data, &faux
	} else if data.CurrentWord == data.TargetWord {
		AttemptLeft = 10
		utils.SaveInFile(structure.Stock{})
		return data, &vrai
	}
	return data, nil
}

func DisplayErrors(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Errors/errors-page.gohtml"))
	_ = tmpl.Execute(w, nil)
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("start/startgame.gohtml"))
		_ = tmpl.Execute(w, nil)
	} else {
		difficulty := r.FormValue("difficulty")
		fmt.Println(difficulty)

		_, _ = manager(nil, &difficulty)
		http.Redirect(w, r, "../", http.StatusSeeOther)
	}
}

func EndgameWin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("finish/endgamewin.gohtml"))
	_ = tmpl.Execute(w, nil)
}

func EndgameLose(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("finish/endgamelose.gohtml"))
	_ = tmpl.Execute(w, nil)
}

func ScoreBoard(r *http.Request) map[string]int {
	username := r.FormValue("username")
	fmt.Println(username)
	score := utils.LoadScoreFile()
	score = map[string]int{}
	fmt.Println(score)
	if username == "" {
		score = map[string]int{username: 0}
	} else {
		score[username] = score[username] + 1
		fmt.Println(score)
	}
	fmt.Println(score)
	utils.SaveScoreInFile(score)
	return score
}

func main() {
	server := http.NewServeMux()
	// url http://localhost:8000/
	server.HandleFunc("/", Home)
	server.HandleFunc("/hangman", Hangman)
	server.HandleFunc("/errors", DisplayErrors)
	server.HandleFunc("/start", StartGame)
	server.HandleFunc("/endwin", EndgameWin)
	server.HandleFunc("/endlose", EndgameLose)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/start")

	http.ListenAndServe(":8000", server)
}
