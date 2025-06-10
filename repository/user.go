package repository

type UserRepository interface {
	RegisterUser(string, string) error
	LoginUser(string, string) error
}