package service

import (
	"database/sql"
	"regexp"
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

func (s *eventService) CreateEvent(newEvent models.Event) error {

	err := s.eventRepo.AddEvent(newEvent)

	if err != nil{
		return httphandle.BadRequest("event is already created")
	}

	return nil
}

func (s *eventService) GetAllEvent() ([]models.Event, error) {

	events,err := s.eventRepo.FetchAllEvent()

	if err != nil{
		return nil,err
	}else if events == nil{
		return nil,httphandle.NotFound("event not found")
	}

	return events,nil
}

func (s *eventService) GetEvent(id string) (*models.Event, error) {

	numberRegex := regexp.MustCompile(`^[0-9]*$`)
	isNumber := numberRegex.MatchString(id)

	if !isNumber {
		return nil,httphandle.BadRequest("id is not a number")
	}
	
	eventId,_ := strconv.Atoi(id)
	event,err := s.eventRepo.FetchEventById(eventId)

	if err == sql.ErrNoRows{
		return nil,httphandle.NotFound("id does not exist")
	}

	return event,nil
}

func (s *eventService) UpdateEvent(name string,description string,id string) error {

	numberRegex := regexp.MustCompile(`^[0-9]*$`)
	isNumber := numberRegex.MatchString(id)

	if !isNumber {
		return httphandle.BadRequest("id is not a number")
	}

	eventId,_ := strconv.Atoi(id)
	err := s.eventRepo.UpdateEventById(name,description,eventId)

	if err == sql.ErrNoRows{
		return httphandle.NotFound("id does not exist")
	}else if err != nil{
		return err
	}

	return nil
}

func (s *eventService) DeleteEvent(id string) (error) {

	numberRegex := regexp.MustCompile(`^[0-9]*$`)
	isNumber := numberRegex.MatchString(id)

	if !isNumber {
		return httphandle.BadRequest("id is not a number")
	}

	eventId,_ := strconv.Atoi(id)

	err := s.eventRepo.DeleteEventById(eventId)

	if err != nil{
		return httphandle.BadRequest(err.Error())
	}

	return nil
}
