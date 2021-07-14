package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

func (u *User) HashPassword(password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(passwordHash)
	return nil
}

const (
	accessTokens  = "fn3014jofp32rkmewnf53poflr"
	refreshTokens = "wfepiernfqenfreornfq"
)

func (u *User) GenToken(userId string) (*AuthTokens, error) {
	acTime, _ := strconv.ParseInt("9000000000000", 10, 64)
	reTime, _ := strconv.ParseInt("86400000000000", 10, 64)
	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Duration(acTime)).Unix()

	td.RtExpires = time.Now().Add(time.Duration(reTime)).Unix()

	var err error

	os.Setenv("ACCESS_SECRET", accessTokens)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: td.AtExpires,
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
	})
	accessToken, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	os.Setenv("REFRESH_SECRET", refreshTokens)
	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: td.AtExpires,
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
	})
	refreshToken, err := token2.SignedString([]byte(os.Getenv("REFRESH_TOKEN")))
	if err != nil {
		return nil, err
	}

	return &AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(time.Duration(td.AtExpires)),
		UserID:       userId,
	}, nil
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) ParseToken(accessToken string) (int, int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("INVALID_SIGNING_METHOD")
		}

		return []byte(accessTokens), nil
	})
	if err != nil {
		return 0, 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, 0, errors.New("CLAIMING_TOKEN_ERROR")
	}
	return claims.UserId, claims.ExpiresAt, nil
}
