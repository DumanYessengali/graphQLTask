package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
	"twoBinPJ/adapters"
	"twoBinPJ/domains/user"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

func GenToken(userId string) (*user.AuthTokens, error) {
	config := adapters.ParseConfig()

	acTime, _ := strconv.ParseInt(config.AtExpires, 10, 64)
	reTime, _ := strconv.ParseInt(config.RtExpires, 10, 64)
	td := &TokenDetails{}

	td.AtExpires = time.Now().Add(time.Duration(acTime)).Unix()

	td.RtExpires = time.Now().Add(time.Duration(reTime)).Unix()

	var err error

	os.Setenv("ACCESS_SECRET", config.AccessTokenKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: td.AtExpires,
		Id:        userId,
		IssuedAt:  time.Now().Unix(),
	})
	accessToken, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	os.Setenv("REFRESH_SECRET", config.RefreshTokenKey)
	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: td.RtExpires,
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
