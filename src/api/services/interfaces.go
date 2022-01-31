package services

import (
	"github.com/amgomez-meli/bi_proxy_server/src/api/domain"
	"github.com/amgomez-meli/bi_proxy_server/src/api/utils"
)

type Reader interface {
	GetTypes() ([]*domain.Proxies_Types, error)
	GetProxyNames() ([]*domain.Proxies_List, error)
	GetTypesAndNames() ([]*domain.Proxies_Names_types, error)
	GetProxyURILISTByType(entityName string) ([]string, error)
	GetUserAgentsRotated() ([]string, error)
	GetUserAgentsUnRotated() ([]string, error)
}

type Repository interface {
	Reader
}

type APIClient interface {
	Get(url string, headers map[string]string, body map[string]interface{}, proxy string) (string, error)
	Post(url string, headers map[string]string, body string, proxy string) (string, error)
}

type Biz_Layer interface {
	ApplyBizConvs(hd utils.Handler, proxy_type string) map[string]string
}
