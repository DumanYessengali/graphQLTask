package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserPostgres(Db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: Db}
}

func (u *UserRepository) GetUserByField(field, value string) (*User, error) {

	var user *User
	row, err := u.DB.Query(fmt.Sprintf("select*from users where %v=$1", field), value)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var u User
		err = row.Scan(&u.Id, &u.Username, &u.Password, &u.Role, &u.Point)

		if err != nil {

			return nil, err
		}

		user = &u
	}
	return user, nil
}

func (u *UserRepository) GetUserByID(id string) (*User, error) {
	return u.GetUserByField("id", id)
}

func (u *UserRepository) GetUserByUsername(username string) (*User, error) {
	return u.GetUserByField("username", username)
}
func (u *UserRepository) GetUserIdByUsername(username string) (string, error) {
	user, err := u.GetUserByField("username", username)
	return user.Id, err
}
func (u *UserRepository) CreateUser(user *User) (*User, error) {
	rows, err := u.DB.Query("insert into users(username, password,role,point) values($1,$2,$3,$4) returning *", user.Username, user.Password, UserRolesUser, 0)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	var mup *User
	for rows.Next() {
		var m User
		err := rows.Scan(&m.Id, &m.Username, &m.Password, &m.Role, &m.Point)
		if err != nil {
			return nil, err
		}
		mup = &m
	}
	_, err = u.DB.Query("insert into user_achievements(user_id, level_achievements_id, created)values($1,$2,$3)", mup.Id, nil, time.Now())

	return mup, nil
}
