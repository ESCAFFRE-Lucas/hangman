package main

import (
	"Correction_hangman/hangman"
	"fmt"
	"math/rand"
	"time"
)

func HangmanManager() {
	draw := hangman.GetDrawFromFile()
	var rightLetters []string
	var wrongLetters []string
	const maxAttempts = 10
	targetWord := hangman.GetRandomWord()
	currentWord := hangman.InitWord(targetWord)
	fmt.Println(targetWord)
	fmt.Println(currentWord)
	for currentWord != targetWord && len(wrongLetters) < maxAttempts {
		userInput := hangman.Input()
		hangman.HandleInput(targetWord, userInput, &currentWord, &rightLetters, &wrongLetters)
		hangman.PrintGame(userInput, currentWord, rightLetters, wrongLetters, draw, maxAttempts)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	HangmanManager()
}
