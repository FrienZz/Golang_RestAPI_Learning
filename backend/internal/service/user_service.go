package service

import (
	"database/sql"
	"regexp"

	"github.com/FrienZz/Golang_RestAPI_Learning/httphandle"
	"github.com/FrienZz/Golang_RestAPI_Learning/internal/port"
	"github.com/FrienZz/Golang_RestAPI_Learning/utils"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo port.UserRepository
}

// Core logic
func NewUserService(userRepo port.UserRepository) port.UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(email string,password string) error {

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	isValid := emailRegex.MatchString(email)

	if !isValid {
		return httphandle.BadRequest("Invalid email format")
	}

	if len(password) < 6{
		return httphandle.BadRequest("Password must be at least 6 characters") 
	}

	hashPassword,err := bcrypt.GenerateFromPassword([]byte(password),14)

	if err != nil{
		return err
	}

	err = s.userRepo.RegisterUser(email,string(hashPassword))

	if err != nil{
		return httphandle.BadRequest("Email already exist")
	}

	return nil
	
}

func (s *userService) LoginUser(email string,password string) (string,error) {

	hashPassword,err := s.userRepo.LoginUser(email,password)

	switch {
	case  err == sql.ErrNoRows:
		return "",httphandle.NotFound("Email does not exist")
	case err != nil:
		return "",err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password))

	if err != nil{
		return "",httphandle.Unauthorized("Password is incorrect. Please try again")
	}

	userId,_ := s.userRepo.GetUserId(email)

	token,err := utils.GenerateToken(email,*userId)

	if err != nil{
		return "",err
	}

	return token,nil
}
