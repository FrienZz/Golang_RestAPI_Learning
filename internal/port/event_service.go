package port

import (
	"github.com/FrienZz/Golang_RestAPI_Learning/internal/models"
)

type EventService interface {
	CreateEvent(models.Event) error
	GetAllEvents() ([]models.Event, error)
	GetEvent(string) (*models.Event, error)
	UpdateEvent(string,string,string,int) error 
	DeleteEvent(string,int) error
}
