package models

type Participant struct {
	Id      int    `json:"id"`
	Pseudo  string `json:"pseudo"`
	Famille int    `json:"famille"`
	Photo   string `json:"photo"`
}
