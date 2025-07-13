package port

import (
	"github.com/FrienZz/Golang_RestAPI_Learning/internal/models"
)

type EventService interface {
	CreateEvent(models.Event,int) error
	GetAllEvents() ([]models.Event, error)
	GetEvent(string) (*models.Event, error)
	UpdateEvent(models.Event,string,int) error 
	DeleteEvent(string,int) error
}
