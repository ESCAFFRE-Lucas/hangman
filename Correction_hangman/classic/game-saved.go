package classic

//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//)
//
//
//const HangmanDataFile = "save.json"
//
//type JsonGame struct { //Create a structure to save the variables and use them after
//	Try          int
//	Letter       string
//	WordProgress string
//	WordToGuess  string
//}
//
//func SaveInFile(data JsonGame, fileName string) { //take and return the json encoding of the structure
//	jsonData, _ := json.Marshal(data) //return the json encoding of data either JsonGame structure
//	err := ioutil.WriteFile(fileName, jsonData, 0644)
//	if err != nil {
//		return
//	}
//	fmt.Println("Game Saved !")
//}
//
//func LoadFile(fileName string) JsonGame { //load the file with the saved param
//	data, err := ioutil.ReadFile(fileName)
//	result := JsonGame{}
//	err = json.Unmarshal(data, &result)
//	if err != nil {
//		return JsonGame{}
//	}
//	return result
//}
//
//func saveJson(){
//	SaveInFile(JsonGame{ //put the structure of the game in the json file stock in the HangmanDataFile variable with the SaveInFile function
//		Try:          attempts,
//		Letter:       stock,
//		WordProgress: hide,
//		WordToGuess:  GetRandomWord(),
//	}, HangmanDataFile)
//}
