package classic

import (
	"math/rand"
	"os"
	"strings"
)

//The function below is used to get all words in a file and stock them in an array of string
func GetWordsFromFile(file *string) []string {
	if file == nil {
		defaultFile := "words.txt"
		file = &defaultFile
	}
	inFile, err := os.ReadFile(*file)
	if err != nil {
		inFile, _ = os.ReadFile("words.txt")
	}
	strFile := string(inFile)
	splitWords := strings.Split(strFile, "\n")
	return splitWords
}

//The function below use the function "GetWordFromFile" and pick a randomword in his array of string (words)
func GetRandomWord(file *string) string {
	words := GetWordsFromFile(file)
	randomWord := words[rand.Intn(len(words))]
	randomWord = strings.Trim(randomWord, "\n\r")
	return randomWord
}

//Use the two functions below  to set a base hidden word with a certain number of letters revealed at the start of a game
func InitWord(word string) string {
	return RevealLetters(word, HideLetters(word))
}

//hide all the letters of a given word
func HideLetters(word string) string {
	var str string
	for i := 0; i < len(word); i++ {
		str += "_"
	}
	return str
}

//take a hidden word and reveal len(word)/2 - 1 letters then return it
func RevealLetters(word, hidden string) string {
	var idx int
	arrHidden := []rune(hidden)
	numberOfLetters := len(word)/2 - 1
	for i := 0; i < numberOfLetters; i++ {
		for {
			idx = rand.Intn(len(word))
			if arrHidden[idx] == rune(word[idx]) {
				continue
			}
			arrHidden[idx] = rune(word[idx])
			break
		}
	}
	return string(arrHidden)
}
