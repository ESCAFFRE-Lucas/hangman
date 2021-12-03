package main

import (
	"Correction_hangman/hangman"
	"fmt"
	"math/rand"
	"time"
)

func HangmanManager() {
	targetWord := hangman.GetRandomWord()
	currentWord := hangman.InitWord(targetWord)
	fmt.Println(targetWord)
	fmt.Println(currentWord)
}

func main() {
	rand.Seed(time.Now().Unix())
	HangmanManager()
}
