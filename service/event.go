package service

import (
	"github.com/FrienZz/Golang_RestAPI_Learning/models"
)

type EventService interface {
	CreateEvent(models.Event) error
	GetAllEvent() ([]models.Event, error)
	GetEvent(int) (*models.Event, error)
	UpdateEvent(string,string,int) error 
	DeleteEvent(int) error
}
