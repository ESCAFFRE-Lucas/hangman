package hangman

func IsNotALetter(str string) bool {
	return str[0] > 'z' || str[0] < 'a'
}
