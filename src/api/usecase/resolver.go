package usecase

type Resolver struct {
	RotatedAgents           []string
	UnRotatedAgents         []string
	RotatedHeaders          []string
	UnrotatedHeaders        []string
	ExcludedDomainsRotated  []string
	ExcludedDomainUnRotated []string
}

func NewResolver(rotAg []string, unrotAg []string, rotH []string, unRotH []string, ExcDomR []string, ExcDomUR []string) *Resolver {
	return &Resolver{
		RotatedAgents:           rotAg,
		UnRotatedAgents:         unrotAg,
		RotatedHeaders:          rotH,
		UnrotatedHeaders:        unRotH,
		ExcludedDomainsRotated:  ExcDomR,
		ExcludedDomainUnRotated: ExcDomUR,
	}
}
