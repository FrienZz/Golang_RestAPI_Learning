package service

import (
	"github.com/FrienZz/Golang_RestAPI_Learning/models"
)

type EventService interface {
	CreateEvent(models.Event) error
	GetAllEvent() ([]models.Event, error)
	GetEvent(string) (*models.Event, error)
	UpdateEvent(string,string,string) error 
	DeleteEvent(string) error
}
