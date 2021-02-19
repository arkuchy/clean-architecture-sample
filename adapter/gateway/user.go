package gateway

import (
	"database/sql"
	"log"

	"github.com/ari1021/clean-architecture/entity"
	"github.com/ari1021/clean-architecture/usecase/port"
)

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) port.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (u *UserRepository) GetUserByID(userID int) (*entity.User, error) {
	row := u.conn.QueryRow("SELECT * FROM `user` WHERE id=?", userID)
	user := entity.User{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
func (u *UserRepository) GetDBConn() *sql.DB {
	return u.conn
}
