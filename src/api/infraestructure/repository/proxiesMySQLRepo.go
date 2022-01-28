package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/amgomez-meli/bi_proxy_server/src/api/config"
	"github.com/amgomez-meli/bi_proxy_server/src/api/domain"
	_ "github.com/go-sql-driver/mysql"
)

type ProxyMySQL struct {
	db *sql.DB
}

func NewProxyRepo() *ProxyMySQL {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err.Error())

	}
	// defer db.Close()

	return &ProxyMySQL{
		db: db,
	}
}

func (r *ProxyMySQL) GetTypes() ([]*domain.Proxies_Types, error) {

	stmt, err := r.db.Prepare(`select id, proxy_type from types`)
	if err != nil {
		return nil, err
	}

	var types []*domain.Proxies_Types
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var z domain.Proxies_Types
		err = rows.Scan(&z.ID, &z.ProxyType)
		if err != nil {
			return nil, err
		}
		types = append(types, &z)

	}

	return types, nil

}

func (r *ProxyMySQL) GetProxyNames() ([]*domain.Proxies_List, error) {

	stmt, err := r.db.Prepare(`select id, service_name from proxy_services`)
	if err != nil {
		return nil, err
	}

	var names []*domain.Proxies_List
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var z domain.Proxies_List
		err = rows.Scan(&z.ID, &z.Proxy_name)
		if err != nil {
			return nil, err
		}
		names = append(names, &z)

	}

	return names, nil

}

func (r *ProxyMySQL) GetTypesAndNames() ([]*domain.Proxies_Names_types, error) {

	stmt, err := r.db.Prepare(`select a.proxy_type, b.proxy_name from type a inner join proxy b on a.ID= b.type`)
	if err != nil {
		return nil, err
	}

	var names []*domain.Proxies_Names_types
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var z domain.Proxies_Names_types
		err = rows.Scan(&z.ProxyType, &z.Proxy_name)
		if err != nil {
			return nil, err
		}
		names = append(names, &z)

	}

	return names, nil

}

func (r *ProxyMySQL) GetProxyURILISTByType(entityName string) ([]string, error) {

	query_str := `select id,proxy_uri from ` + entityName

	stmt, err := r.db.Prepare(query_str)
	if err != nil {
		return nil, err
	}

	var uris []string
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var z domain.Proxies_URIS
		err = rows.Scan(&z.ID, &z.Proxy_url)
		if err != nil {
			return nil, err
		}
		uris = append(uris, z.Proxy_url)

	}

	return uris, nil

}

func (r *ProxyMySQL) GetUserAgentsRotated() ([]string, error) {

	var a []string
	return a, nil

}
func (r *ProxyMySQL) GetUserAgentsUnRotated() ([]string, error) {

	var a []string
	return a, nil

}
