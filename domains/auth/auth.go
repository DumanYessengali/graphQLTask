package auth

import (
	"context"
	"github.com/jmoiron/sqlx"
	"twoBinPJ/domains/user"
)

type IAuthService interface {
	SignIn(ctx context.Context, username, password string) (*user.AuthTokens, *user.User, error)
	SignUp(ctx context.Context, username, password string) (string, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*user.AuthTokens, *user.User, error)
	Logout(ctx context.Context, refreshToken string) (string, error)
}
type IAuthRepository interface {
	FillTheAuth(id, rToken string) error
	DeleteAuth(token string) error
	CheckTokenBeforeRefresh(userId string, token string) (bool, error)
	UpdateAuth(userid, token string) error
	CheckIfExistsAuth(userId string) (bool, error)
	CheckTokenIfExist(id string, token string) (bool, error)
}

type AuthModule struct {
	IAuthService
}

type Auth struct {
	ID       string
	RefToken string
	UserID   string
}

func NewAuthModule(Db *sqlx.DB) *AuthModule {
	userRepository := user.NewUserPostgres(Db)
	userService := user.NewUserService(userRepository)
	authRepository := NewAuthPostgres(Db)
	return &AuthModule{
		IAuthService: NewAuthService(userService, authRepository),
	}
}
