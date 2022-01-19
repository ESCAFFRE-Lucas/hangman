package utils

import (
	"encoding/json"
	"hangman_web/structure"
	"io/ioutil"
)

const HangmanDataFile = "save.json"

const ScoreBoardDataFile = "score.json"

func SaveInFile(data structure.Stock) { //take and return the json encoding of the structure
	jsonData, _ := json.Marshal(data) //return the json encoding of data either structure.Stock structure
	err := ioutil.WriteFile(HangmanDataFile, jsonData, 0644)
	if err != nil {
		return
	}
}

func LoadFile() structure.Stock { //load the file with the saved param
	data, err := ioutil.ReadFile(HangmanDataFile)
	result := structure.Stock{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return structure.Stock{}
	}
	return result
}

func SaveScoreInFile(data map[string]int) { //take and return the json encoding of the structure
	jsonData, _ := json.Marshal(data) //return the json encoding of data either structure.Stock structure
	err := ioutil.WriteFile(ScoreBoardDataFile, jsonData, 0644)
	if err != nil {
		return
	}
}

func LoadScoreFile() map[string]int { //load the file with the saved param
	data, err := ioutil.ReadFile(ScoreBoardDataFile)
	result := make(map[string]int)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return make(map[string]int)
	}
	return result
}
