package port

type UserRepository interface {
	RegisterUser(string, string) error
	LoginUser(string, string) (string, error)
	GetUserId(string) (*int, error)
}