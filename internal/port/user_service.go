package port

type UserService interface {
	RegisterUser(string, string) error
	LoginUser(string, string) (string, error)
}
