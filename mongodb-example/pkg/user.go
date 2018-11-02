package root

type User struct {
	Id       string `json:"id"`
	Username string `json: "username"`
	Password string `json: "username"`
}

type UserService interface {
	CreateUser(u *User) error
	GetByUsername(username string) (*User, error)
}
