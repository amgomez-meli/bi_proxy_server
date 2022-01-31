package usecase

type Entity_Handler interface {
	GetConfigurationAsMap() map[string]bool
	GetURI() string
	GetHeadersAsMap() map[string]string
	GetBody() map[string]interface{}
}

type Biz_Conf interface {
	AddHeaders() map[string]string
}
