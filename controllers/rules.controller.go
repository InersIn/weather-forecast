package controllers

import (
	"fmt"
	"strconv"
	"weather-forecast/models"
)

func ForwardChaining(json_map []models.AnswerResponse) int {
	// var Rules models.Answer
	var status int = 1
	r1 := json_map[1].Data["ans"] == "1" || json_map[2].Data["ans"] == "1"
	r2 := json_map[0].Data["ans"] == "0" && r1

	r3 := json_map[5].Data["ans"] == "1"
	r4 := -99
	if r3 {
		n4, _ := strconv.ParseInt(json_map[4].Data["ans"], 6, 12)
		n5, _ := strconv.ParseInt(json_map[6].Data["ans"], 6, 12)
		r4 = int(n4 - n5)
		fmt.Println("N:", n4, n5, r4)
	}

	fmt.Println(r3, json_map[5].Data["ans"], r4)
	if r2 && r4 == 0 {
		status = 0
	} else if r1 && json_map[5].Data["ans"] == "0" {
		status = 2
	}
	fmt.Println(r1, json_map[5].Data["ans"], r4)

	fmt.Println(status)
	return status
}
