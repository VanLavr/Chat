package models

type Response struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}
