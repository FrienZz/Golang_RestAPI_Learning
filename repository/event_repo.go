package repository

import (
	"database/sql"
	"errors"

	"github.com/FrienZz/Golang_RestAPI_Learning/models"
)

type eventRepositoryDB struct {
	db *sql.DB
}

//Adapter Database 
func NewEventRepositoryDB(db *sql.DB) EventRepository{
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

func (r *eventRepositoryDB) FetchAllEvent() ([]models.Event,error){
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

func (r *eventRepositoryDB) UpdateEventById(name string,description string,id int) error{


	updateData,err := r.FetchEventById(id)

	if err != nil{
		return err
	}

	stmt := "UPDATE events SET name = $1,description = $2 WHERE id = $3"
	
	if name != ""{
		updateData.Name = name
	}

	if description != ""{
		updateData.Description = description
	}
	
	_,err = r.db.Exec(stmt,updateData.Name,updateData.Description,id)

	if err != nil{
		return err
	}

	return nil
}


func (r *eventRepositoryDB) DeleteEventById(id int) (error){
	stmt := "DELETE FROM events WHERE id = $1"
	result,err := r.db.Exec(stmt,id)

	if err != nil{
		return err
	}
	
	rowAffected,err := result.RowsAffected()

	if err != nil{
		return err
	}

	if rowAffected == 0{
		return errors.New("id does not exist")
	}

	return nil

}