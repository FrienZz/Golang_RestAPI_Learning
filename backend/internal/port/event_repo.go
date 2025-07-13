package port

import "github.com/FrienZz/Golang_RestAPI_Learning/internal/models"

//Secondary Port
type EventRepository interface{
	AddEvent(models.Event,int)  (error)
	FetchAllEvents()  ([]models.Event,error)
	FetchEventById(int) (*models.Event,error)
	UpdateEventById(models.Event,int,int) (error)
	DeleteEventById(int,int)  (error)
}