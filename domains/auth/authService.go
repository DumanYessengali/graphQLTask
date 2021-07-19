package auth

import (
	"context"
	"errors"
	"log"
	"twoBinPJ/domains/user"
)

type AuthService struct {
	Serv user.IUserService
	Repo IAuthRepository
}

func NewAuthService(serv *user.UserService, repo IAuthRepository) *AuthService {
	return &AuthService{
		Serv: serv,
		Repo: repo,
	}
}

func (a *AuthService) SignIn(ctx context.Context, username, password string) (*user.AuthTokens, *user.User, error) {
	user, err := a.Serv.GetUserByUsernameService(username)
	if err != nil {
		log.Printf("user does not exist: %s", err)
		return nil, nil, errors.New("INITIALIZING_USER_ERROR")
	}
	err = user.ComparePassword(password)
	if err != nil {
		log.Printf("error while comparing passwords: %s", err)
		return nil, nil, errors.New("COMPARING_PASSWORD_ERROR")
	}

	token, err := GenToken(user.Id)
	if err != nil {
		log.Printf("creating token error: %s", err)
		return nil, nil, errors.New("INITIALIZING_TOKENS_ERROR")
	}
	userIdCheck, err := a.Repo.CheckIfExistsAuth(user.Id)
	if err != nil {
		log.Printf("error while check user: %s", err)
		return nil, nil, errors.New("INITIALIZING_USER_ERROR")
	}
	if !userIdCheck {
		err = a.Repo.FillTheAuth(user.Id, token.RefreshToken)
		if err != nil {
			log.Printf("error while inserting new row on auth table: %s", err)
			return nil, nil, errors.New("INITIALIZING_NEW_TOKEN_ERROR")
		}
	} else {
		err = a.Repo.UpdateAuth(user.Id, token.RefreshToken)
		if err != nil {
			log.Printf("error while inserting new row on auth table: %s", err)
			return nil, nil, errors.New("UPDATING_TOKEN_ERROR")
		}
	}
	return token, user, nil
}

func (a *AuthService) SignUp(ctx context.Context, username, password string) (string, error) {
	user := &user.User{
		Username: username,
	}
	err := user.HashPassword(password)
	if err != nil {
		log.Printf("error while creating password, and %s", err)
		return "", errors.New("INITIALIZING_PASSWORD_ERROR")
	}

	if _, err := a.Serv.CreateUserService(user); err != nil {
		log.Printf("error while creating new user: %s", err)
		return "", errors.New("INITIALIZING_USER_ERROR")
	}
	message := "you are successfully sing up"
	return message, nil
}

func (a *AuthService) RefreshTokens(ctx context.Context, refreshToken string) (*user.AuthTokens, *user.User, error) {
	currentUser, err := user.ForContext(ctx)

	if err != nil {
		log.Printf("token is incorrect or wrong1: %s", err)
		return nil, nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	checkToken, err := a.Repo.CheckTokenBeforeRefresh(currentUser.Id, refreshToken)
	if err != nil {
		log.Printf("token is incorrect or wrong2: %s", err)
		return nil, nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if checkToken {
		token, err := GenToken(currentUser.Id)
		if err != nil {
			log.Printf("creating token error: %s", err)
			return nil, nil, errors.New("INITIALIZING_TOKENS_ERROR")
		}
		err = a.Repo.UpdateAuth(currentUser.Id, token.RefreshToken)
		if err != nil {
			log.Printf("error while inserting new row on auth table: %s", err)
			return nil, nil, errors.New("UPDATING_TOKEN_ERROR")
		}
		return token, currentUser, nil
	} else {
		log.Printf("refreshing token is incorrect or wrong")
		return nil, nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
}

func (a *AuthService) Logout(ctx context.Context, refreshToken string) (string, error) {
	currentUser, err := user.ForContext(ctx)

	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return "", errors.New("INITIALIZING_TOKEN_ERROR")
	}
	checkToken, err := a.Repo.CheckTokenBeforeRefresh(currentUser.Id, refreshToken)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return "", errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if checkToken {
		err = a.Repo.DeleteAuth(refreshToken)
		if err != nil {
			log.Printf("error while deleting user from auth table: %s", err)
			return "", errors.New("DELETING_USER_ERROR")
		}
		message := "you are logout"
		return message, nil
	} else {
		message := "you are not logout"
		return message, nil
	}
}
