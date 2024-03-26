package models

import "time"

type Greet struct {
	Id              string    `json:"id"`
	GreetingContent string    `json:"greet_content"`
	CreatedAt       time.Time `json:"created_at"`
	Lenguage        string    `json:"lenguage"`
	UserId          string    `json:"user_id"`
}
