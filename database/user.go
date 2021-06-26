package database

import (
	"database/sql"
	"log"
	"reflect"

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

		err := ScanEntity(rows, &u)
		if err != nil {
			return users, err
		}

		users = append(users, u)
	}

	return users, nil
}

func ScanEntity(row Row, modelEntity interface{}) error {

	typeOf := reflect.ValueOf(modelEntity).Elem()
	kindOf := typeOf.Kind()
	fieldsNumber := typeOf.NumField()

	structFields := make([]interface{}, fieldsNumber)

	if kindOf != reflect.Struct {
		log.Fatal("unexpected type")
	}

	for i := 0; i < fieldsNumber; i++ {
		structFields[i] = typeOf.Field(i).Addr().Interface()
	}

	err := row.Scan(structFields...)
	if err != nil {
		return err
	}

	return nil
}
