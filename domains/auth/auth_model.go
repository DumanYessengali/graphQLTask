package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
	"twoBinPJ/domains/user"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

const (
	refreshTokens = "wfepiernfqenfreornfq"
	accessTokens  = "fn3014jofp32rkmewnf53poflr"
)

func GenToken(userId string) (*user.AuthTokens, error) {
	acTime, _ := strconv.ParseInt("9000000000000", 10, 64)
	reTime, _ := strconv.ParseInt("86400000000000", 10, 64)
	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Duration(acTime)).Unix()

	td.RtExpires = time.Now().Add(time.Duration(reTime)).Unix()

	var err error

	os.Setenv("ACCESS_SECRET", accessTokens)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: td.AtExpires,
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
	})
	accessToken, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	os.Setenv("REFRESH_SECRET", refreshTokens)
	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: td.AtExpires,
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
	})
	refreshToken, err := token2.SignedString([]byte(os.Getenv("REFRESH_TOKEN")))
	if err != nil {
		return nil, err
	}

	return &user.AuthTokens{
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

func ParseToken(accessToken string) (int, int64, error) {
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
