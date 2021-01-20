package datamysql

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/go-sql-driver/mysql"
)

var (
	instance   *Connection
	dbInstance *sql.DB
	once       sync.Once
	connError  error
)

type Connection struct {
	db *sql.DB
}

func getMySQLConfig() *mysql.Config {
	config := mysql.NewConfig()

	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%s", 0, "port")
	config.DBName = ""
	config.User = ""
	config.Passwd = ""
	config.ParseTime = true

	return config
}

func GetDB() (*sql.DB, error) {
	once.Do(func() {
		mysqlconfig := getMySQLConfig()

		db, err := sql.Open("mysql", mysqlconfig.FormatDSN())
		if err != nil {
			connError = err
			return
		}

		dbInstance = db
	})

	return dbInstance, connError
}
