package service

import (
	"database/sql"
	"strconv"

	"github.com/FrienZz/Golang_RestAPI_Learning/httphandle"
	"github.com/FrienZz/Golang_RestAPI_Learning/internal/models"
	"github.com/FrienZz/Golang_RestAPI_Learning/internal/port"
)

type eventService struct {
	eventRepo port.EventRepository
}
//Core logic
func NewEventService(eventRepo port.EventRepository) port.EventService{
	return &eventService{eventRepo: eventRepo}
}

func (s *eventService) CreateEvent(newEvent models.Event,userId int) error {

	err := s.eventRepo.AddEvent(newEvent,userId)

	if err != nil{
		return httphandle.BadRequest("Event is already created")
	}

	return nil
}

func (s *eventService) GetAllEvents() ([]models.Event, error) {

	events,err := s.eventRepo.FetchAllEvents()

	switch {
	case  err != nil:
		return nil,err
	case events == nil:
		return nil,httphandle.NotFound("Event not found")
	}

	return events,nil
}

func (s *eventService) GetEvent(id string) (*models.Event, error) {
	
	eventId,err := strconv.Atoi(id)

	if err != nil{
		return nil,httphandle.BadRequest("Id is not a number")
	}
	
	event,err := s.eventRepo.FetchEventById(eventId)

	switch {
	case  err == sql.ErrNoRows:
		return nil,httphandle.NotFound("Id does not exist")
	case err != nil:
		return nil,err
	}
		
	return event,nil
}

func (s *eventService) UpdateEvent(title string,description string,id string,userId int) error {

	eventId,err := strconv.Atoi(id)

	if err != nil{
		return httphandle.BadRequest("Id is not a number")
	}

	err = s.eventRepo.UpdateEventById(title,description,eventId,userId)

	switch{
	case err == sql.ErrNoRows:
		return httphandle.NotFound("Id does not exist")
	case err != nil && err.Error() == "unauthorized access":
		return httphandle.Unauthorized("Unauthorized Access")
	case err != nil:
		return err
	}
	

	return nil
}

func (s *eventService) DeleteEvent(id string,userId int) (error) {

	eventId,err := strconv.Atoi(id)

	if err != nil{
		return httphandle.BadRequest("Id is not a number")
	}

	err = s.eventRepo.DeleteEventById(eventId,userId)

	switch{
	case err == sql.ErrNoRows:
		return httphandle.NotFound("Id does not exist")
	case err != nil && err.Error() == "unauthorized access":
		return httphandle.Unauthorized("Unauthorized Access")
	case err != nil:
		return err
	}

	return nil
}
