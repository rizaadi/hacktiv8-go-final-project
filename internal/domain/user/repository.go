package user

type UserRepository interface {
	CreateUser(user *User) (err error)
	GetUserByEmail(email string) (user User, err error)
}
