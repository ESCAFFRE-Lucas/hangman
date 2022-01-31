package classic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Rentrez une lettre svp : ")
	for scanner.Scan() {
		if len(scanner.Text()) > 1 || IsNotALetter(scanner.Text()) {
			fmt.Println("Mets qu'une seule lettre !!")
			continue
		}
		break
	}
	return scanner.Text()
}

//What the function below do : if a lowercase letter is detected, first we search if the letter is in "word", then reveal
//the same letter in "hiddenword" then put it in "right" if the letter is in "word", else we put it in "wrong" and tell the
//player to choose a lowercase letter. After that we decrease the number of attempt by the len of "wrong"
func HandleInput(word, letter string, hiddenWord *string, right, wrong *[]string, attempts int) {
	if LowercaseOnly(letter) == true {
		//1er partie cherchez si la letter est dans le mot ou non
		res := strings.Index(word, letter)
		//2eme partie Si oui, la mettre dans le mot à l'index prévu. Si non, ne rien faire.
		if res != -1 {
			arr := []rune(*hiddenWord)
			for i := 0; i < len(word); i++ {
				if word[i] == letter[0] {
					arr[i] = rune(letter[0])
				}
			}
			*hiddenWord = string(arr)
			if !contains(*right, letter) {
				*right = append(*right, letter)
			}
		} else {
			if !contains(*wrong, letter) {
				*wrong = append(*wrong, letter)
			}
		}
		attempts = attempts - len(*wrong)
	} else {
		fmt.Println("Please, choose a lowercase letter !")
	}
}

func LowercaseOnly(letter string) bool { //This function is used to know if a letter is lowercase or not (if yes it return True, else False)
	if letter < "a" || letter > "z" {
		return false
	}
	return true
}
