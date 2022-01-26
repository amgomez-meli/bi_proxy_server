package main

import (
	"fmt"

	"github.com/amgomez-meli/bi_proxy_server/src/api/infraestructure/repository"
	"github.com/amgomez-meli/bi_proxy_server/src/api/services"
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

}
