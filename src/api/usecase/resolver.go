package usecase

type Resolver struct {
	MapOfConfigs map[string][]string
}

func NewResolver(rotAg []string, unrotAg []string, rotH []string, unRotH []string, ExcDomR []string, ExcDomUR []string) *Resolver {

	map_conf := make(map[string][]string)
	map_conf["RotatedAgents"] = rotAg
	map_conf["UnRotatedAgents"] = unrotAg
	map_conf["RotatedHeaders"] = rotH
	map_conf["UnrotatedHeaders"] = unRotH
	map_conf["ExcludedDomainsRotated"] = ExcDomR
	map_conf["ExcludedDomainUnRotated"] = ExcDomUR

	return &Resolver{
		MapOfConfigs: map_conf,
	}
}

func (t *Resolver) ApplyBizConvs(l Entity_Handler, proxy_type string, proxy_uri string) {

}
