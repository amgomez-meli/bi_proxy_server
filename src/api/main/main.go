package main

import (
	"fmt"

	"github.com/amgomez-meli/bi_proxy_server/src/api/infraestructure/repository"
	"github.com/amgomez-meli/bi_proxy_server/src/api/services"
	"github.com/amgomez-meli/bi_proxy_server/src/api/utils"
)

func main() {

	fmt.Println("starting ...")

	repo := repository.NewProxyRepo()

	fmt.Println("starting 2 ...")

	output := services.NewLoadConfigService(repo)
	fmt.Println("starting 3 ...")
	output.LoadProxiesList()
	t := output.GetURIFromProvider("shader")

	for _, a := range t {
		fmt.Println("new imput: " + a)

	}
	jsonstring := `
	{
		"domain":"https://shopee.com.mx/",
		"priority": 1,
		"config": { "oxylabs": true, "crawlera_from_brasil": false, "crawlera_from_mexicod": false,"crawlera": false,"shader_chile": false, "shader": false },
		"options": {
			"url": "https://shopee.com.mx/api/v4/shop/get_shop_detail?shopid=521171866",
			"method": "GET",
			"headers":{
				"testheader": "hellotest",
				"cataz": "wwwww",
				"vat": "eeeee"
			}
			
	}
			}
    `
	parser := utils.GenericParserMap(jsonstring)
	EntityHandler := utils.ProcessJsonMsg(parser)
	maps := EntityHandler.GetConfigurationAsMap()

	for k, v := range maps {

		fmt.Printf("AAAA", k, v)
	}

	url := EntityHandler.GetURI()

	fmt.Println("THE URI IS: ", url)

	header := EntityHandler.GetHeadersAsMap()["testheader"]
	header1 := EntityHandler.GetHeadersAsMap()["vat"]
	fmt.Println("THE header  IS: ", header)
	fmt.Println("THE other  header  IS: ", header1)

}
