package utils

import (
	"encoding/json"
	"fmt"
)

type MapParser struct {
	JSONAsMap map[string]interface{}
}

func GenericParserMap(jsonString string) *MapParser {

	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {

		fmt.Println(err)
	}

	return &MapParser{

		JSONAsMap: result,
	}

}

func (z *MapParser) GetStringField(key string) string {

	return z.JSONAsMap[key].(string)

}

func (z *MapParser) GetIntField(key string) int {

	return z.JSONAsMap[key].(int)

}

func (z *MapParser) GetNestedStringField(object string, key string) string {
	field := ""
	genObj := z.JSONAsMap[object].(map[string]interface{})

	for key1, value := range genObj {
		if key1 == key {
			field = value.(string)

		}

	}
	return field

}

func (z *MapParser) GetNestedIntField(object string, key string) int {
	var field int
	genObj := z.JSONAsMap[object].(map[string]interface{})

	for key1, value := range genObj {
		if key1 == key {
			field = value.(int)

		}

	}
	return field

}

func (z *MapParser) GetObjAsmAP(object string) map[string]interface{} {

	return z.JSONAsMap[object].(map[string]interface{})

}
