package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Questions struct {
	Root  map[string]QuestionData `json:"root"`
	Chain map[string]QuestionData `json:"chain"`
}

type QuestionData struct {
	Id     string            `json:"id"`
	Type   string            `json:"type"`
	Desc   string            `json:"desc"`
	Labels map[string]string `json:"label"`
	Ans    map[string]string `json:"ans"`
}

var res Questions
var err error

func _Init() {
	jsonFile, err := os.Open("./database/rules.json")
	if err != nil {
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &res)
}

func GetQuestions(types string, id string) (data QuestionData, err error) {
	_Init()

	if types == "root" {
		data = res.Root[id]
	} else if types == "chain" {
		data = res.Chain[id]
	}
	data.Id = id

	if data.Labels == nil || data.Ans == nil {
		err = errors.New("data not found")
	}
	return
}
