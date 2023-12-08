package postgres

import (
	"errors"
	"fmt"
	"time"

	"github.com/Murolando/hakaton_final_api/ent"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) SignUp(user ent.User) (int64, error) {
	var id int64
	t := time.Now()
	dateTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	user.RegistrationDateTime = &dateTime
	query := fmt.Sprintf(`
	INSERT INTO "%s"
	(password_hash,email,role_id, registration_datetime, age) 
	values ($1 , $2 ,$3 ,$4, $5) RETURNING id`, userTable)
	row := r.db.QueryRow(query,
		user.PasswordHash, user.Login, user.RoleId, user.RegistrationDateTime,user.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r AuthPostgres) GetUserByLoginAndPassword(mail *string, password *string) (int64, error) {
	var id int64
	query := fmt.Sprintf(`SELECT id FROM "%s" WHERE login = $1`, userTable)
	row := r.db.QueryRow(query, mail)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	query = fmt.Sprintf(`SELECT id FROM "%s" WHERE login = $1 AND password_hash = $2`, userTable)
	row = r.db.QueryRow(query, mail, password)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r AuthPostgres) SetSession(userId int64, refresh string, expiredAt string) error {
	var id int
	query := fmt.Sprintf(`
	UPDATE "%s" SET refresh = $1,expired_at =$2
	WHERE id = $3
	RETURNING id`, userTable)
	row := r.db.QueryRow(query, refresh, expiredAt, userId)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (r AuthPostgres) GetByRefreshToken(refresh string) (int64, error) {
	var id int64
	var expiredAt time.Time
	query := fmt.Sprintf(`SELECT id,expired_at FROM "%s" WHERE refresh = $1`, userTable)
	row := r.db.QueryRow(query, refresh)
	if err := row.Scan(&id, &expiredAt); err != nil {
		return 0, errors.New("bad token")
	}
	if time.Now().Sub(expiredAt) > 0 {
		return 0, errors.New("refresh token expired")
	}
	return id, nil
}
