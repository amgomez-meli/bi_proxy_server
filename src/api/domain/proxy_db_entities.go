package domain

type Proxies_Types struct {
	ID        int
	ProxyType string
}

type Proxies_List struct {
	ID         int
	Proxy_name string
	Proxy_type int
}

type Proxies_URIS struct {
	ID        int
	Proxy_url string
}

type Proxies_Names_types struct {
	Proxy_name string
	ProxyType  string
}

type Proxies_user_agents struct {
	Agent string
}
