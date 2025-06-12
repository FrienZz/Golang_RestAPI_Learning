package service

import (
	"regexp"

	"github.com/FrienZz/Golang_RestAPI_Learning/httphandle"
	"github.com/FrienZz/Golang_RestAPI_Learning/repository"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

// Core logic
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(email string,password string) error {

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	isValid := emailRegex.MatchString(email)

	if !isValid {
		return httphandle.BadRequest("invalid email format")
	}

	if len(password) < 6{
		return httphandle.BadRequest("password must be at least 6 characters") 
	}

	hashPassword,err := bcrypt.GenerateFromPassword([]byte(password),14)

	if err != nil{
		return err
	}

	err = s.userRepo.RegisterUser(email,string(hashPassword))

	if err != nil{
		return httphandle.BadRequest("email already exist")
	}

	return nil
	
}

func (s *userService) LoginUser(email string,password string) error {


	hashPassword,err := s.userRepo.LoginUser(email,password)

	if err != nil{

		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password))

	if err != nil{
		return httphandle.NotFound("password is incorrect. Please try again")
	}

	return nil
}
