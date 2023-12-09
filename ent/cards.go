package ent

type Card struct {
	Id          int    `json:"card-id"`
	Word        string `json:"word"`
	Description string `json:"description"`
}
