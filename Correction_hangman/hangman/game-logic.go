package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Rentrez une lettre svp : ")
	for scanner.Scan() {
		if len(scanner.Text()) > 1 || IsNotALetter(scanner.Text()) {
			fmt.Println("Bâtard mets qu'une seule lettre je t'ai dit ! Tu sais pas lire ?!")
			continue
		}
		fmt.Println(scanner.Text())
		break
	}
	return scanner.Text()
}
