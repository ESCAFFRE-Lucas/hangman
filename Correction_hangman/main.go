package main

import (
	"Correction_hangman/hangman"
	"fmt"
	"math/rand"
	"time"
)

func HangmanManager() {
	//var rightLetter []string
	//var wrongLetter []string
	targetWord := hangman.GetRandomWord()
	currentWord := hangman.InitWord(targetWord)
	fmt.Println(targetWord)
	fmt.Println(currentWord)
}

func main() {
	rand.Seed(time.Now().Unix())
	//HangmanManager()
	hangman.Input()
}
