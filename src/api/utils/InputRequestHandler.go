package utils

import (
	"fmt"
	"reflect"
)

type Handler struct {
	Parser Json_Parser
}

func ProcessJsonMsg(p Json_Parser) *Handler {

	return &Handler{
		Parser: p,
	}

}

func (z *Handler) GetConfigurationAsMap() map[string]bool {
	outputmap := make(map[string]bool)
	a := z.Parser.GetObjAsmAP("config")
	for k, v := range a {

		outputmap[k] = v.(bool)

	}

	return outputmap

}

func (z *Handler) GetURI() string {
	obj := z.Parser.GetObjAsmAP("options")

	return obj["url"].(string)

}

func (z *Handler) GetHeadersAsMap() map[string]string {

	outputmap := make(map[string]string)
	obj := z.Parser.GetObjAsmAP("options")["headers"]

	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			v := val.MapIndex(e)
			switch t := v.Interface().(type) {

			case string:
				outputmap[e.String()] = t
				fmt.Println(e, t)

			default:
				fmt.Println("not found")

			}

		}

	}

	return outputmap

}

func (z *Handler) GetBody() map[string]interface{} {

	obj := z.Parser.GetObjAsmAP("options")["body"]
	return obj.(map[string]interface{})

}
