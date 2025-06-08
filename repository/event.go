package repository

import "github.com/FrienZz/Golang_RestAPI_Learning/models"

//Secondary Port
type EventRepository interface{
	AddEvent(models.Event)  (error)
	FetchAllEvent()  ([]models.Event,error)
	FetchEventById(int) (*models.Event,error)
	UpdateEventById(string,string,int) (error)
	DeleteEventById(int)  (error)
}