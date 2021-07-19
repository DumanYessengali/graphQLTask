package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"twoBinPJ/adapters"
	"twoBinPJ/domains/user"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header
			id, err := ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			user := user.User{Id: id}

			user.Id = id
			ctx := context.WithValue(r.Context(), "user", &user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ParseToken(tokenStr string) (string, error) {
	config := adapters.ParseConfig()

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AccessTokenKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["jti"].(string)
		return id, nil
	} else {
		return "", err
	}
}
