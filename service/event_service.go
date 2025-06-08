package service

import (
	"github.com/FrienZz/Golang_RestAPI_Learning/models"
	"github.com/FrienZz/Golang_RestAPI_Learning/repository"
)

type eventService struct {
	eventRepo repository.EventRepository
}
//Core logic
func NewEventService(eventRepo repository.EventRepository) EventService{
	return &eventService{eventRepo: eventRepo}
}

func (s *eventService) CreateEvent(newEvent models.Event) error {

	err := s.eventRepo.AddEvent(newEvent)

	if err != nil{
		return err
	}

	return nil
}

func (s *eventService) GetAllEvent() ([]models.Event, error) {

	events,err := s.eventRepo.FetchAllEvent()

	if err != nil{
		return nil,err
	}

	
	return events,nil
}

func (s *eventService) GetEvent(id int) (*models.Event, error) {

	event,err := s.eventRepo.FetchEventById(id)

	if err != nil{
		return nil,err
	}

	return event,nil
}

func (s *eventService) UpdateEvent(name string,description string,id int) error {

	err := s.eventRepo.UpdateEventById(name,description,id)

	if err != nil{
		return err
	}

	return nil
}

func (s *eventService) DeleteEvent(id int) (error) {

	err := s.eventRepo.DeleteEventById(id)

	if err != nil{
		return err
	}

	return nil
}
