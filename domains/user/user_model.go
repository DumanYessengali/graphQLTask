package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) HashPassword(password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(passwordHash)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func ForContext(ctx context.Context) (*User, error) {
	raw, ok := ctx.Value("user").(*User)
	if !ok {
		return nil, errors.New("no user in context")
	}
	return raw, nil
}
