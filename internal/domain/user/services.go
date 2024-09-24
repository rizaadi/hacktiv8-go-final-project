package user

type service struct {
	repo UserRepository
}

// GetUserByEmail implements Service.
func (s *service) GetUserByEmail(email string) (user User, err error) {
	return s.repo.GetUserByEmail(email)
}

// CreateUser implements Service.
func (s *service) CreateUser(user *User) (err error) {
	return s.repo.CreateUser(user)
}

type Service interface {
	CreateUser(user *User) (err error)
	GetUserByEmail(email string) (user User, err error)
}

func NewService(r UserRepository) Service {
	return &service{repo: r}
}
