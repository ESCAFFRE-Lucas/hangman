package classic

//The function below is used to know if the string "arr" (that will be a letter for us)
//is in the array of string "arr" (that will be a word for us)
func contains(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}
