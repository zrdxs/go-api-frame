package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/MarceloZardoBR/go-api-frame/domain/interfaces"
	"github.com/MarceloZardoBR/go-api-frame/infra/config"
	_ "github.com/go-sql-driver/mysql"
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

type Row interface {
	Scan(...interface{}) error
}

type RowScanner interface {
	ScanRow(Row, interface{}) error
}

func Instance(cfg *config.Config) (interfaces.Data, error) {

	database := &Database{}

	db, err := StartDB(cfg)
	if err != nil {
		return database, err
	}

	database.db = db
	database.userRepo = &userRepo{db}

	return database, err
}

// CreateConfiguration returns db conn string
func createConfiguration(cfg *config.Config) (psqlconn string) {

	host := cfg.DBHost
	port := cfg.DBPort
	user := cfg.DBUser
	password := cfg.DBPassword
	dbname := cfg.DBName

	switch cfg.DBService {
	case "postgres":
		psqlconn = fmt.Sprintf(" port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", port, host, user, password, dbname)
	case "mysql":
		psqlconn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	default:
		return psqlconn
	}

	return psqlconn
}

// StartDB connection
func StartDB(cfg *config.Config) (*sql.DB, error) {
	once.Do(func() {
		log.Println("Starting DB Connection...")
		psqlconn := createConfiguration(cfg)
		db, err := sql.Open(cfg.DBService, psqlconn)
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
