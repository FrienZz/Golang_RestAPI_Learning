package service

type UserService interface {
	RegisterUser(string, string) error
	LoginUser(string, string) error
}
