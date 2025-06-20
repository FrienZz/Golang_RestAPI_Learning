package port

import "github.com/FrienZz/Golang_RestAPI_Learning/internal/models"

//Secondary Port
type EventRepository interface{
	AddEvent(models.Event)  (error)
	FetchAllEvents()  ([]models.Event,error)
	FetchEventById(int) (*models.Event,error)
	UpdateEventById(string,string,int,int) (error)
	DeleteEventById(int)  (error)
}