package services

import (
	"math/rand"
	"time"

	"github.com/amgomez-meli/bi_proxy_server/src/api/domain"
)

type ProxyServ struct {
	Client      APIClient
	Proxy_list  map[string][]string
	Proxy_types []*domain.Proxies_Names_types
}

func NewProxyService(a APIClient, proxys_lists map[string][]string, proxys_types []*domain.Proxies_Names_types) *ProxyServ {

	return &ProxyServ{
		Client:      a,
		Proxy_list:  proxys_lists,
		Proxy_types: proxys_types,
	}

}

func (z *ProxyServ) AssignProxy(s map[string]bool) string {
	var proxy_type string
	var proxy_name string
	var proxy_uri string

	for _, currtype := range z.Proxy_types {

		if s[currtype.Proxy_name] {
			proxy_type = currtype.ProxyType
			proxy_name = currtype.Proxy_name

		}

	}

	if proxy_type == "1" {
		proxy_uri = z.Proxy_list[proxy_name][0]

	} else {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(z.Proxy_list) + 1)
		proxy_uri = z.Proxy_list[proxy_name][n]

	}

	return proxy_uri

}
