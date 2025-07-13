package models

import "time"

type Event struct {
	Event_id    int 		`json:"id"`
	Title        string 		`json:"name"`
	Description string 		`json:"description"`
	Created_at  time.Time 	`json:"created_at"`
	User_id     int 		`json:"user_id"`
}