package models

import "time"

type User struct {
	User_id    	int       	`json:"id"`
	Email       string    	`json:"email"`
	Password 	string    	`json:"password"`
	Created_at  time.Time 	`json:"created_at"`
}