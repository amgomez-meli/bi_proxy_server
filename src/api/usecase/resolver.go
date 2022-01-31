package usecase

import "fmt"

type Resolver struct {
	MapOfConfigs map[string][]string
	RotHead      map[string]string
	UnRotHead    map[string]string
}

func NewResolver(rotH map[string]string, unRotHmap map[string]string, ExcDomR []string, ExcDomUR []string) *Resolver {

	map_conf := make(map[string][]string)
	map_conf["ExcludedDomainsRotated"] = ExcDomR
	map_conf["ExcludedDomainUnRotated"] = ExcDomUR

	return &Resolver{
		MapOfConfigs: map_conf,
		RotHead:      rotH,
		UnRotHead:    unRotHmap,
	}
}

func (t *Resolver) ApplyBizConvs(l Entity_Handler, proxy_type string, agent string) map[string]string {

	headers := l.GetHeadersAsMap()

	switch proxy_type {

	case "1":

		headers = RotatedHeadAdd(headers, t.RotHead, agent)

		fmt.Println("www")

	case "2":
		headers = UnRotatedHeadAdd(headers, t.UnRotHead, agent)
		fmt.Println("www")

	case "3":
		headers = APIHeadAdd(headers)
		fmt.Println("www")

	default:
		fmt.Println("zaw")

	}

	return headers

}
