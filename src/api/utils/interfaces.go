package utils

type Json_Parser interface {
	GetStringField(key string) string
	GetIntField(key string) int
	GetNestedStringField(object string, key string) string
	GetNestedIntField(object string, key string) int
	GetObjAsmAP(object string) map[string]interface{}
}
