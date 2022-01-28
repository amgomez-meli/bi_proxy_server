package services

import "fmt"

type Proxyconfig struct {
	Configurations  map[string][]string
	RotatedAgents   []string
	UnRotatedAgents []string
	repo            Repository
}

func NewLoadConfigService(r Repository) *Proxyconfig {
	configmap := make(map[string][]string)

	return &Proxyconfig{
		Configurations: configmap,
		repo:           r,
	}

}

func (z *Proxyconfig) LoadProxiesList() {

	names, err := z.repo.GetProxyNames()

	if err != nil {
		fmt.Println(err)

		return
	}

	for _, name := range names {
		fmt.Println(name.Proxy_name)
		proxy_List, err := z.repo.GetProxyURILISTByType(name.Proxy_name)

		if err != nil {
			fmt.Println(err)

			return
		}

		z.Configurations[name.Proxy_name] = proxy_List

	}

}

func (z *Proxyconfig) LoadUserAgents() {
	rotated, err := z.repo.GetUserAgentsRotated()

	unrotated, err := z.repo.GetUserAgentsUnRotated()

	if err != nil {
		fmt.Println(err)

	}
	z.RotatedAgents = rotated
	z.UnRotatedAgents = unrotated

}

func (z *Proxyconfig) GetURIFromProvider(keytosearch string) []string {

	return z.Configurations[keytosearch]

}

func (z *Proxyconfig) GetAgents() ([]string, []string) {

	return z.RotatedAgents, z.UnRotatedAgents

}
