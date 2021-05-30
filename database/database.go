package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/MarceloZardoBR/go-api-frame/domain/interfaces"
	_ "github.com/lib/pq"
)

var (
	once     sync.Once
	instance *sql.DB
	dberr    error
)

type Database struct {
	db *sql.DB

	userRepo *userRepo
}

func Instance() (interfaces.Data, error) {

	database := &Database{}

	db, err := StartDB()
	if err != nil {
		return database, err
	}

	database.db = db
	database.userRepo = &userRepo{db}

	return database, err
}

// CreateConfiguration returns db conn string
func createConfiguration() string {

	host := "localhost"
	port := 5432
	user := "postgres"
	password := ""
	dbname := "db_teste"

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return psqlconn
}

// StartDB connection
func StartDB() (*sql.DB, error) {
	once.Do(func() {
		log.Println("Starting DB Connection...")
		psqlconn := createConfiguration()
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			dberr = err
			return
		}
		instance = db
	})

	return instance, dberr
}

func (d *Database) UserRepo() interfaces.UserRepo {
	return d.userRepo
}
