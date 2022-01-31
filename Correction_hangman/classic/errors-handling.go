package classic

func IsNotALetter(str string) bool { //This function is used to know if a letter is lowercase or not (if yes it return True, else False)
	return str[0] > 'z' || str[0] < 'a'
}
