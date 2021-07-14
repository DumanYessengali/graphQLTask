package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"twoBinPJ/apps/api1/graph/generated"
	"twoBinPJ/apps/api1/graph/model"
	"twoBinPJ/domains/user"
)

func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInUser) (*model.AuthResponse, error) {
	token, user, err := r.AuthModule.SignIn(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AuthTokens: token,
		User:       user,
	}, nil
}

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpUser) (*model.Message, error) {
	message, err := r.AuthModule.SignUp(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: message}, nil
}

func (r *mutationResolver) RefreshTokens(ctx context.Context, input model.Refresh) (*model.AuthResponse, error) {
	token, user, err := r.AuthModule.RefreshTokens(ctx, input.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AuthTokens: token,
		User:       user,
	}, nil
}

func (r *mutationResolver) Logout(ctx context.Context, input model.Refresh) (*model.Message, error) {
	message, err := r.AuthModule.Logout(ctx, input.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: message}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	return r.UserModule.GetUserByIDService(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
