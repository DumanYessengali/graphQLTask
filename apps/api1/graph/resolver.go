package graph

import (
	"twoBinPJ/domains/auth"
	"twoBinPJ/domains/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	AuthModule auth.IAuthService
	UserModule user.IUserService
}