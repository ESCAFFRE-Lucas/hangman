package hangman

import (
	"fmt"
	"os"
	"strings"
)

func PrintGame(input, word string, right, wrong, draw []string, attempts int) {
	fmt.Println(input)
	fmt.Println(word)
	fmt.Println(right, wrong)
	fmt.Println(draw[len(wrong)-1])
	fmt.Println("Nombre d'essais restants : ", attempts-len(wrong))
	if attempts-len(wrong) == 0 {
		fmt.Println("Arrivederci")
	}
}

func GetDrawFromFile() []string {
	draw, _ := os.ReadFile("hangman.txt")
	strDraw := string(draw)
	splitDraws := strings.Split(strDraw, "=========")
	for i := 0; i < len(splitDraws)-1; i++ {
		splitDraws[i] += "=========\n"
	}
	return splitDraws
}
