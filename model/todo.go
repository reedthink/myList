package model

//Todo_model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
	Email  string `json:"email"`
}