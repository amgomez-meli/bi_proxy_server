package utils

type Handler struct {
	Parser Json_Parser
}

func ProcessJsonMsg(p Json_Parser) *Handler {

	return &Handler{
		Parser: p,
	}

}

func (z *Handler) GetConfigurationAsMap() map[string]bool {
	var outputmap map[string]bool
	a := z.Parser.GetObjAsmAP("config")
	for k, v := range a {
		outputmap[k] = v.(bool)

	}

	return outputmap

}
