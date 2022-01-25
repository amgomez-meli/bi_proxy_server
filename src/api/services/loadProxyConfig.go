package services

type Proxyconfig struct {
	Configurations map[string][]string
	repo           Repository
}

func newLoadConfigService(r Repository) *Proxyconfig {
	configmap := make(map[string][]string)

	return &Proxyconfig{
		Configurations: configmap,
		repo:           r,
	}

}

func (z *Proxyconfig) loadProxiesList() {

	names, err := z.repo.GetProxyNames()

	if err != nil {

		return
	}

	for _, name := range names {
		proxy_List, err := z.repo.GetProxyURILISTByType(name.Proxy_name)

		if err != nil {

			return
		}

		z.Configurations[name.Proxy_name] = proxy_List

	}

}
