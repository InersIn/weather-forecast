package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AnswerResponse struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

type Error struct {
	Status  string `json:status`
	Message string `json:message`
}
