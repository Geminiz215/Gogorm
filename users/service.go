package users

type Service interface {
	FindOne(username UserRequestFindOne) (User, error)
	CreateUser(user UserRequestCreate) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s service) FindOne(userRequest UserRequestFindOne) (User, error) {
	username := userRequest.Username
	user, err := s.repository.FindOne(username)

	return user, err
}

func (s service) CreateUser(userRequest UserRequestCreate) (User, error) {

	user := User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	newuser, err := s.repository.CreateUser(user)

	return newuser, err
}
