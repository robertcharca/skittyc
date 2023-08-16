package kfeatures

import (
	"errors"
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
)

var empty = errors.New("Empty values not allowed")

func ChangingValues(attribute, value , section string) error {
	

	if attribute != "" && value != "" && section != "" {
		chAttribute := strings.ReplaceAll(attribute, " ", "_")

		var chValue string
		chValue = chAttribute + " " + value	

		if !kittyc.ModifyingAtLine(chAttribute, chValue) {
			kittyc.WritingAtLine(section, chValue)
		}

		return nil
	}

	return empty 
}
