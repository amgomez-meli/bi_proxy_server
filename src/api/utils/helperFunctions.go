package utils

func MergeStringMaps(map1 map[string]string, map2 map[string]string) map[string]string {

	for k, v := range map2 {
		map1[k] = v

	}

	return map1

}

func AddMapItemIfNotExists(z map[string]string, keyToEval string, valuetoadd string) map[string]string {

	if _, ok := z[keyToEval]; !ok {

		z[keyToEval] = valuetoadd

	}

	return z

}
