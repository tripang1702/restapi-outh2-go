package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(server string, port string, user string, password string, database string) *sqlx.DB {
	var db *sqlx.DB
	var err error

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, server, port, database)
	db, err = sqlx.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	return db

}
