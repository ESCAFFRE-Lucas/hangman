package main

import (
	"classic"
	"fmt"
	"hangman_web/structure"
	"hangman_web/utils"
	"html/template"
	"net/http"
)

// Home This function below permit to execute the index template (the main page of the game)
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
	//if classic.IsNotALetter(letter) {
	//	http.Redirect(w, r, "/errors", http.StatusSeeOther)
	//} else {
	//	http.Redirect(w, r, "../", http.StatusSeeOther)
	//}
}

// Redirect This function below permit to redirect to another page with an url
func Redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusSeeOther)
}

//This function below is the most important, it uses the stock structure and some normal hangman functions
//to play the hangman game on the server, while saving the progress. if hidden word == current word then win,
//if attempts < 10 then lose (if the len of the input is >= 2, then attempts down by 2)
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

// DisplayErrors This function below permit to execute the errors template, wich will handle when the player type a non lowercase letter
func DisplayErrors(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Errors/errors-page.gohtml"))
	_ = tmpl.Execute(w, nil)
}

// StartGame This function below execute the start template, wich will ask the username and the difficulty of the player
func StartGame(w http.ResponseWriter, r *http.Request) {
	//username := r.FormValue("username")
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

// EndgameWin This function below execute the endgamewin template, wich will redirect to a "You Won !" page
func EndgameWin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("finish/endgamewin.gohtml"))
	_ = tmpl.Execute(w, nil)
}

// EndgameLose This function below execute the endgamelose template, wich will redirect to a "You Lost !" page
func EndgameLose(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("finish/endgamelose.gohtml"))
	_ = tmpl.Execute(w, nil)
}

//The function below show a scoreboard in the main page, with the player's username and his score (+1 per win)
//func ScoreBoard(user *string) map[string]int {
//	fmt.Println(*username)
//	score := utils.LoadScoreFile()
//	score = map[string]int{}
//	fmt.Println(score)
//	if *username == "" {
//		score = map[string]int{*username: 0}
//	} else {
//		score[*username] = score[*username] + 1
//		fmt.Println(score)
//	}
//	fmt.Println(score)
//	utils.SaveScoreInFile(score)
//	return score
//}

//This function below start the server and handle some functions to begin to play the game
func main() {
	server := http.NewServeMux()
	// url http://localhost:8000/
	server.HandleFunc("/", Home)
	server.HandleFunc("/hangman", Hangman)
	server.HandleFunc("/errors", DisplayErrors)
	server.HandleFunc("/start", StartGame)
	//server.HandleFunc("/endwin", EndgameWin)
	//server.HandleFunc("/endlose", EndgameLose)

	server.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// listen to the port 8000
	fmt.Println("server listening on http://localhost:8000/start")

	_ = http.ListenAndServe(":8000", server)

}
