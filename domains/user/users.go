package user

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type IUserRepository interface {
	GetUserByField(field, value string) (*User, error)
	GetUserByID(id string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	CreateUser(user *User) (*User, error)
}

type IUserService interface {
	GetUserByIDService(id string) (*User, error)
	GetUserByUsernameService(username string) (*User, error)
	CreateUserService(user *User) (*User, error)
}

type UserModule struct {
	IUserService
}

type User struct {
	Id       string
	Username string
	Password string
	Role     UserRoles
}
type UserRoles int

const (
	UserRolesUser      UserRoles = iota // 0
	UserRolesModerator                  // 1
	UserRolesAdmin                      // 2
	UserRolesManager                    // 3
)

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    time.Time
	UserID       string
}

func NewUserModule(Db *sqlx.DB) *UserModule {
	userRepository := NewUserPostgres(Db)
	return &UserModule{
		IUserService: NewUserService(userRepository),
	}
}
