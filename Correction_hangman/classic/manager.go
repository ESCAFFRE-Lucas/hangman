package classic

//func HangmanManager() {
//	draw := game_progress.GetDrawFromFile()
//	var rightLetters []string
//	var wrongLetters []string
//	const maxAttempts = 10
//	targetWord := game_progress.GetRandomWord()
//	currentWord := game_progress.InitWord(targetWord)
//	fmt.Println(currentWord)
//	for currentWord != targetWord && len(wrongLetters) < maxAttempts {
//		userInput := game_progress.Input()
//		game_progress.HandleInput(targetWord, userInput, &currentWord, &rightLetters, &wrongLetters)
//		game_progress.PrintGame(currentWord, rightLetters, wrongLetters, draw, maxAttempts)
//	}
//}
//
//func main() {
//	rand.Seed(time.Now().Unix())
//	HangmanManager()
//}
