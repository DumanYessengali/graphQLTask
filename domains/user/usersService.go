package user

type UserService struct {
	Repo IUserRepository
}

func NewUserService(repos IUserRepository) *UserService {
	return &UserService{Repo: repos}
}

func (s *UserService) GetUserByIDService(id string) (*User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *UserService) GetUserByUsernameService(username string) (*User, error) {
	return s.Repo.GetUserByUsername(username)
}

func (s *UserService) CreateUserService(user *User) (*User, error) {
	return s.Repo.CreateUser(user)
}
