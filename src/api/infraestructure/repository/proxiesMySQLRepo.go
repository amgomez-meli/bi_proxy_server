package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/amgomez-meli/bi_proxy_server/src/api/config"
	"github.com/amgomez-meli/bi_proxy_server/src/api/domain"
)

type ProxyMySQL struct {
	db *sql.DB
}

func NewClientMySQL(db *sql.DB) *ProxyMySQL {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())

	}
	defer db.Close()

	return &ProxyMySQL{
		db: db,
	}
}

func (r *ProxyMySQL) getTypes(*domain.Proxy_name, error) {

	// stmt, err := r.db.Prepare(`select id, proxy_type from types`)

}
