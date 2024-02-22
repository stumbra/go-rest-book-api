package models

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Published   int    `json:"published"`
	InStore     bool   `json:"in_store"`
}
