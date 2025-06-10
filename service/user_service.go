package service

import (
	"github.com/FrienZz/Golang_RestAPI_Learning/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

// Core logic
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(email string,password string) error {


	err := s.userRepo.RegisterUser(email,password)

	if err != nil{
		return err
	}

	return nil
	
}

func (s *userService) LoginUser(email string,password string) error {

	err := s.userRepo.LoginUser(email,password)

	if err != nil{
		return err
	}

	return nil
}
