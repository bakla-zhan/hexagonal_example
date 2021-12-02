package dbuser

import (
	"database/sql"
	"hexagonal/register"
)

type UserMapper struct {
	db *sql.DB
}

func (um *UserMapper) SaveUser(u register.User) error {
	_, err := um.db.Exec("")
	if err != nil {
		return err
	}
	return nil
}
