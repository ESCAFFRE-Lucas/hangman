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

func HandleInput(word, letter string, hiddenWord *string, right, wrong *[]string, attempts *int) {
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
		//Attempts(letter, attempts, wrong)
	}
}

//func Attempts(letter string, attempts *int, wrong *[]string) {
//	count := 1
//	 if len(letter) == 1 {
//		 *attempts = 10 - len(*wrong)
//	 } else {
//		 count += 2
//		 *attempts -= 2
//	 }
//}
