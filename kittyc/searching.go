package kittyc

import "errors"

var notFound = errors.New("Value not found in array")

func SearchingValue(list []string, value string) (bool, error) {
	for _, sv := range list {
		if sv == value {
			return true, nil
		} 
	}

	return false, notFound
}

func SearchingSimilarValues(list []string, svalue string) (bool, error) {	
	for l := 0; l <= len(list); l++ {
		if list[l][:4] == svalue[:4] {
			return true, nil
		}
	}
	return false, notFound
}

func ConvertStringToList(s string) []string {	
	var sList []string

	list := []rune(s)

	for _, v := range list { 
		sList = append(sList, string(v))
	}

	return sList
}
