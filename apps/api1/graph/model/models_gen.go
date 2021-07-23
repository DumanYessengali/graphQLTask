// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
	"twoBinPJ/domains/user"
)

type AuthResponse struct {
	AuthTokens *user.AuthTokens `json:"authTokens"`
	User       *user.User       `json:"user"`
}

type CreateProject struct {
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	Description      string `json:"description"`
}

type CreateReport struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Comments    string `json:"Comments"`
	Seriousness string `json:"Seriousness"`
}

type CreateVulnerability struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Message struct {
	Message string `json:"message"`
}

type Refresh struct {
	RefreshToken string `json:"refreshToken"`
}

type SignInUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateProject struct {
	Name             *string `json:"name"`
	ShortDescription *string `json:"shortDescription"`
	Description      *string `json:"description"`
}

type UpdateReport struct {
	Name        *string `json:"Name"`
	Description *string `json:"Description"`
	Comments    *string `json:"Comments"`
	Seriousness *string `json:"Seriousness"`
}

type UpdateVulnerability struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type LevelAchievements struct {
	ID                  string    `json:"id"`
	UserID              int       `json:"userId"`
	LevelAchievementsID int       `json:"levelAchievementsId"`
	Created             time.Time `json:"created"`
}
