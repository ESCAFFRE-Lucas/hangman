package classic

import (
	"math/rand"
	"os"
	"strings"
)

func GetWordsFromFile() []string {
	inFile, _ := os.ReadFile("words.txt")
	strFile := string(inFile)
	splitWords := strings.Split(strFile, "\n")
	return splitWords
}

func GetRandomWord() string {
	words := GetWordsFromFile()
	randomWord := words[rand.Intn(len(words))]
	randomWord = strings.Trim(randomWord, "\n\r")
	return randomWord
}

func InitWord(word string) string {
	return RevealLetters(word, HideLetters(word))
}

func HideLetters(word string) string {
	var str string
	for i := 0; i < len(word); i++ {
		str += "_ "
	}
	return str
}

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
