package domain

type Input_Request struct {
	Domain   string                `json:"domain"`
	Priority int                   `json:"priority"`
	Config   Input_Request_Config  `json:"config"`
	Options  Input_Request_Options `json:"options"`
}

type Input_Request_Config struct {
	Oxylabs              bool
	Crawlera_from_brasil bool
	Crawlera_from_mexico bool
	Crawlera             bool
	Shader               bool
	Shader_chile         bool
}

type Input_Request_Options struct {
	Url     string
	Method  string
	Headers map[string]bool
	Body    map[string]bool
}
