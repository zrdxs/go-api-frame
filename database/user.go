package database

import (
	"database/sql"

	"github.com/MarceloZardoBR/go-api-frame/domain/entity"
)

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) GetAll() ([]entity.User, error) {
	var users []entity.User

	query := `select user_id, name, age, email from users;`

	rows, err := u.db.Query(query)
	if err != nil {
		return users, err
	}

	defer rows.Close()
	for rows.Next() {
		u := entity.User{}

		err = ScanRow(rows, &u)
		if err != nil {
			return users, err
		}

		users = append(users, u)
	}

	return users, nil
}

func ScanRow(r Row, user *entity.User) (err error) {
	err = r.Scan(&user.UserID,
		&user.Name,
		&user.Age,
		&user.Email,
	)
	if err != nil {
		return err
	}

	return nil
}
