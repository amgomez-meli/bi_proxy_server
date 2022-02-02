package infraestructure

import (
	"context"
	"fmt"

	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
)

func Get(uri string, headers map[string]string, body string) {

	resp := rest.Get("https://api.restfulsite.com/resource", rest.Context(context.Background()))

	fmt.Println(resp)

}
