package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	generated1 "twoBinPJ/apps/api1/graph/generated"
	"twoBinPJ/apps/api1/graph/model"
	"twoBinPJ/middleware"
)

func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInUser) (*model.AuthResponse, error) {
	user, err := r.UserService.GetUserByUsernameService(input.Username)

	if err != nil {
		log.Printf("user does not exist: %s", err)
		return nil, errors.New("INITIALIZING_USER_ERROR")
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		log.Printf("error while comparing passwords: %s", err)
		return nil, errors.New("COMPARING_PASSWORD_ERROR")
	}

	token, err := user.GenToken(user.ID)
	if err != nil {
		log.Printf("creating token error: %s", err)
		return nil, errors.New("INITIALIZING_TOKENS_ERROR")
	}
	userIdCheck, err := r.UserService.CheckIfExistsAuthService(user.ID)
	if err != nil {
		log.Printf("error while check user: %s", err)
		return nil, errors.New("INITIALIZING_USER_ERROR")
	}
	if !userIdCheck {
		err = r.UserService.FillTheAuthService(user.ID, token.RefreshToken)
		if err != nil {
			log.Printf("error while inserting new row on auth table: %s", err)
			return nil, errors.New("INITIALIZING_NEW_TOKEN_ERROR")
		}
	} else {
		err = r.UserService.UpdateAuthService(user.ID, token.RefreshToken)
		if err != nil {
			log.Printf("error while inserting new row on auth table: %s", err)
			return nil, errors.New("UPDATING_TOKEN_ERROR")
		}
	}
	return &model.AuthResponse{
		AuthTokens: token,
		User:       user,
	}, nil
}

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpUser) (*model.Message, error) {
	user := &model.User{
		Username: input.Username,
	}
	err := user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while creating password, and %s", err)
		return nil, errors.New("INITIALIZING_PASSWORD_ERROR")
	}

	if _, err := r.UserService.CreateUserService(user); err != nil {
		log.Printf("error while creating new user: %s", err)
		return nil, errors.New("INITIALIZING_USER_ERROR")
	}

	return &model.Message{Message: "you are successfully sing up"}, nil
}

func (r *mutationResolver) RefreshTokens(ctx context.Context, input model.Refresh) (*model.AuthResponse, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	checkToken, err := r.UserService.CheckTokenBeforeRefreshService(currentUser.ID, input.RefreshToken)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if checkToken {
		token, err := currentUser.GenToken(currentUser.ID)
		if err != nil {
			log.Printf("creating token error: %s", err)
			return nil, errors.New("INITIALIZING_TOKENS_ERROR")
		}

		err = r.UserService.UpdateAuthService(currentUser.ID, token.RefreshToken)
		if err != nil {
			log.Printf("error while inserting new row on auth table: %s", err)
			return nil, errors.New("UPDATING_TOKEN_ERROR")
		}
		return &model.AuthResponse{
			AuthTokens: token,
			User:       currentUser,
		}, nil
	} else {
		log.Printf("refreshing token is incorrect or wrong")
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
}

func (r *mutationResolver) Logout(ctx context.Context, input model.Refresh) (*model.Message, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)

	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	checkToken, err := r.UserService.CheckTokenBeforeRefreshService(currentUser.ID, input.RefreshToken)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if checkToken {
		err = r.UserService.DeleteAuthService(input.RefreshToken)
		if err != nil {
			log.Printf("error while deleting user from auth table: %s", err)
			return nil, errors.New("DELETING_USER_ERROR")
		}
		return &model.Message{Message: "you are logout"}, nil
	} else {
		return &model.Message{Message: "you are not logout"}, errors.New("LOGOUT_ERROR")
	}
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UserService.GetUserByIDService(id)
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
