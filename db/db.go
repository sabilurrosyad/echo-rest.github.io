package db

import (
	"database/sql"
	"myapp/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString error")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
	DB = db
}

func CreateCon() *sql.DB {
	return DB
}
