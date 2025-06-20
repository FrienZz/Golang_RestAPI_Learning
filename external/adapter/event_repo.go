package adapter

import (
	"database/sql"
	"errors"

	"github.com/FrienZz/Golang_RestAPI_Learning/internal/models"
	"github.com/FrienZz/Golang_RestAPI_Learning/internal/port"
)

type eventRepositoryDB struct {
	db *sql.DB
}

//Adapter Database 
func NewEventRepositoryDB(db *sql.DB) port.EventRepository{
	return &eventRepositoryDB{db : db}
}

func (r *eventRepositoryDB) AddEvent(newEvent models.Event) (error){
	stmt := "INSERT INTO events(name,description,user_id) VALUES($1,$2,$3)"
	_,err := r.db.Exec(stmt,newEvent.Name,newEvent.Description,newEvent.User_id)

	if err != nil{
		return err
	}
	
	return nil
}

func (r *eventRepositoryDB) FetchAllEvents() ([]models.Event,error){
	query := "SELECT * FROM events ORDER BY id ASC"
	rows,err := r.db.Query(query)
	if err != nil{
		return nil,err
	}

	var events []models.Event
	for rows.Next() {
		var event models.Event 
		err = rows.Scan(&event.Event_id,&event.Name,&event.Description,&event.Created_at,&event.User_id)
	
		if err != nil{
			return nil,err
		}
		
		events = append(events,event)
	}
	return events,nil

}

func (r *eventRepositoryDB) FetchEventById(id int) (*models.Event,error){
	query := "SELECT * FROM events WHERE id = $1 "
	row := r.db.QueryRow(query,id)

	var event models.Event 
	err := row.Scan(&event.Event_id,&event.Name,&event.Description,&event.Created_at,&event.User_id)
	if err != nil {
		return nil,err
	}

	return &event,nil
}

func (r *eventRepositoryDB) UpdateEventById(name string,description string,eventId int,userId int) error{


	updateData,err := r.FetchEventById(eventId)

	if err != nil{
		return err
	}

	stmt := "UPDATE events SET name = $1,description = $2 WHERE id = $3 AND user_id = $4"
	
	if name != ""{
		updateData.Name = name
	}

	if description != ""{
		updateData.Description = description
	}
	
	_,err = r.db.Exec(stmt,updateData.Name,updateData.Description,eventId,userId)

	if err != nil{
		return err
	}

	if updateData.User_id != userId{
		return errors.New("unauthorized access")
	}

	return nil
}


func (r *eventRepositoryDB) DeleteEventById(id int,userId int) (error){

	eventData,err := r.FetchEventById(id)

	if err != nil{
		return err
	}

	stmt := "DELETE FROM events WHERE id = $1 AND user_id =$2"
	_,err = r.db.Exec(stmt,id,userId)

	if err != nil{
		return err
	}
	
	if eventData.User_id != userId{
		return errors.New("unauthorized access")
	}

	return nil

}