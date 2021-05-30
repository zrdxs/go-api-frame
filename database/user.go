package database

import (
	"database/sql"

	"github.com/MarceloZardoBR/go-api-frame/domain/entity"
)

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) GetAll() []entity.User {

	return nil
}
