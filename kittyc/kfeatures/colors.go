package kfeatures

import (
	"fmt"
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
)

func SetColors(path string) {
	colors, err := kittyc.GettingMultipleValues(path, "color")
	if err != nil {
		fmt.Println(err.Error())
	}	

	var colorKeys []string

	for _, color := range colors {
		parts := strings.Fields(color)
		if len(parts) == 2 {
			colorKeys = append(colorKeys, parts[0])
		} else {
			break
		}
	}

	fmt.Println(colorKeys)

	errValues := ChangingMultipleValues(colorKeys, colors, "# Colors")
	if errValues != nil {
		fmt.Println(err.Error())
	}
}
