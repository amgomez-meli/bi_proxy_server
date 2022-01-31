package usecase

import "github.com/amgomez-meli/bi_proxy_server/src/api/utils"

func RotatedHeadAdd(head map[string]string, genericHeadmap map[string]string, agent string) map[string]string {

	z := utils.MergeStringMaps(head, genericHeadmap)
	z = utils.AddMapItemIfNotExists(z, "user-agent", agent)

	return z
}

func UnRotatedHeadAdd(head map[string]string, genericHeadmap map[string]string, agent string) map[string]string {

	z := utils.MergeStringMaps(head, genericHeadmap)
	z = utils.AddMapItemIfNotExists(z, "user-agent", agent)

	return z

}

func APIHeadAdd(head map[string]string) map[string]string {
	return head

}
