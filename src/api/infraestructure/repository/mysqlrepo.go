package repository

import (
	"database/sql"
	"fmt"
	"log"

	"src/api/config"
)

type UserMySQL struct {
	db *sql.DB
}

func NewClientMySQL(db *sql.DB) *UserMySQL {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())

	}
	defer db.Close()

	return &UserMySQL{
		db: db,
	}
}
