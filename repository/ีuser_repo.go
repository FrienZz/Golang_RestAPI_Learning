package repository

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type userRepositoryDB struct {
	db *sql.DB
}

// Adapter Database
func NewUserRepositoryDB(db *sql.DB) UserRepository {
	return &userRepositoryDB{db: db}
}

func (r *userRepositoryDB) RegisterUser(email string,password string) error {

	hashPassword,err := bcrypt.GenerateFromPassword([]byte(password),14)

	if err != nil {
		return err
	}

	stmt := "INSERT INTO users(email,password) VALUES($1,$2)"
	_,err = r.db.Exec(stmt,email,hashPassword)

	if err != nil{
		return err
	}

	return nil
}

func (r *userRepositoryDB) LoginUser(email string,password string) error {
	
	query := "SELECT password FROM users WHERE email = $1"
	result := r.db.QueryRow(query,email)

	var hashPassword string
	err := result.Scan(&hashPassword)

	if err != nil{
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password))

	if err != nil{
		return err
	}

	return nil

}
