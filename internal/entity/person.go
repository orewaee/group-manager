package entity

import "time"

type Person struct {
	Id        Id        `json:"id"`
	GroupId   Id        `json:"group_id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Birthday  time.Time `json:"birthday"`
}
