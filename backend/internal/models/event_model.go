package models

import "time"

type Event struct {
	Event_id    int 		`json:"id"`
	Title        string 	`json:"title"`
	Description string 		`json:"description"`
	ImageUrl    string 		`json:"image_url"`
	EventDate   time.Time	`json:"event_date"`
	Location	string	    `json:"location"`
	Created_at  time.Time 	`json:"created_at"`
	User_id     int 		`json:"user_id"`
}