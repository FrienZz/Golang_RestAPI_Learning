package adapter

import (
	"database/sql"

	"github.com/FrienZz/Golang_RestAPI_Learning/internal/port"
)

type userRepositoryDB struct {
	db *sql.DB
}

// Adapter Database
func NewUserRepositoryDB(db *sql.DB) port.UserRepository {
	return &userRepositoryDB{db: db}
}

func (r *userRepositoryDB) RegisterUser(email string,hashPassword string) error {

	stmt := "INSERT INTO users(email,password) VALUES($1,$2)"
	_,err := r.db.Exec(stmt,email,hashPassword)

	if err != nil{
		return err
	}

	return nil
}

func (r *userRepositoryDB) LoginUser(email string,password string) (string,error) {
	
	query := "SELECT password FROM users WHERE email = $1"
	result := r.db.QueryRow(query,email)

	var hashPassword string
	err := result.Scan(&hashPassword)
	
	if err != nil{
		return "",err
	}

	return hashPassword,nil

}
