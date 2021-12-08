package hangman

import "fmt"

func PrintGame(input, word string, right, wrong []string, attempts int) {
	fmt.Println(input)
	fmt.Println(word)
	fmt.Println(right, wrong)
	fmt.Println("Nombre d'essais restants : ", attempts-len(wrong))
}
