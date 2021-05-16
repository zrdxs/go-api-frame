package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	once     sync.Once
	instance *sql.DB
	dberr    error
)

// CreateConfiguration returns db conn string
func CreateConfiguration(host string, port int64, user string, password string, dbname string) string {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return psqlconn
}

// StartDB connection
func StartDB(psqlconn string) (*sql.DB, error) {
	once.Do(func() {
		log.Println("Starting DB Connection...")
		db, err := sql.Open("postgres", psqlconn)
		if err != nil {
			dberr = err
			return
		}
		instance = db
	})

	return instance, dberr
}
