package services

import "github.com/amgomez-meli/bi_proxy_server/src/api/domain"

type Reader interface {
	GetTypes() ([]*domain.Proxies_Types, error)
	GetProxyNames() ([]*domain.Proxies_List, error)
	GetTypesAndNames() ([]*domain.Proxies_Names_types, error)
	GetProxyURILISTByType(entityName string) ([]string, error)
}

type Repository interface {
	Reader
}
