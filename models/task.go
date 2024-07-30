package models

type Task struct {
	Id        int    `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

