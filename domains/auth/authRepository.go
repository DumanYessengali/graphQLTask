package auth

import "github.com/jmoiron/sqlx"

type AuthRepository struct {
	DB *sqlx.DB
}

func NewAuthPostgres(Db *sqlx.DB) *AuthRepository {
	return &AuthRepository{DB: Db}
}
func (u *AuthRepository) CheckTokenIfExist(id string, token string) (bool, error) {
	row, err := u.DB.Query("select*from auth where user_id=$1", id)

	if err != nil {
		return false, err
	}
	var mup *Auth
	for row.Next() {
		var m Auth
		err := row.Scan(&m.ID, &m.RefToken, &m.UserID)
		if err != nil {
			return false, err
		}

		mup = &m
	}
	if mup.RefToken == token {
		return true, nil
	} else {
		return false, nil
	}
}
func (u *AuthRepository) CheckIfExistsAuth(userId string) (bool, error) {
	var ch bool
	query := "select exists(select * from auth where user_id=$1)"
	row := u.DB.QueryRow(query, userId)
	if err := row.Scan(&ch); err != nil {
		return false, err
	}
	return ch, nil
}

func (u *AuthRepository) UpdateAuth(userid, token string) error {

	_, err := u.DB.Query("update auth set refresh_token=$1 where user_id=$2 ", token, userid)
	return err
}

func (u *AuthRepository) CheckTokenBeforeRefresh(userId string, token string) (bool, error) {
	var refresh_token string
	query := "select refresh_token from auth where user_id=$1"
	row := u.DB.QueryRow(query, userId)
	if err := row.Scan(&refresh_token); err != nil {
		return false, err
	}
	if refresh_token == token {
		return true, nil
	}
	return false, nil
}

func (u *AuthRepository) DeleteAuth(token string) error {
	query := "Delete from auth where refresh_token=$1"
	_, err := u.DB.Exec(query, token)
	if err != nil {
		return err
	}
	return nil
}

func (u *AuthRepository) FillTheAuth(id, rToken string) error {
	_, err := u.DB.Query("insert into auth(user_id, refresh_token) values($1,$2) returning *", id, rToken)
	return err
}
