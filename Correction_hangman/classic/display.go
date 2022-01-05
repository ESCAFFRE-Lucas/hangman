package classic

import (
	"fmt"
	"os"
	"strings"
)

func PrintGame(word string, right, wrong, draw []string, attempts int) {
	fmt.Println("Right letters :", right, "\n", "Wrong letters :", wrong)
	fmt.Println(draw[len(wrong)])
	fmt.Println("Nombre d'essais restants : ", attempts-len(wrong))
	fmt.Println(word)
	if attempts-len(wrong) == 0 {
		fmt.Println("Arrivederci")
	} else if len(right) == len(word) {
		fmt.Println("Good Job !!")
	}
}

func GetDrawFromFile() []string {
	draw, _ := os.ReadFile("game-progress-classic.txt")
	strDraw := string(draw)
	splitDraws := strings.Split(strDraw, "=========")
	for i := 0; i < len(splitDraws)-1; i++ {
		splitDraws[i] += "=========\n"
	}
	return splitDraws
}
