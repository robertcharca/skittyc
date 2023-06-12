package kittyc

import "errors"

var notFound = errors.New("Value not found in array")

func SearchingValue (list []string, value string) (bool, error) {
	for _, sv := range list {
		if sv == value {
			return true, nil
		} 
	}

	return false, notFound
}
