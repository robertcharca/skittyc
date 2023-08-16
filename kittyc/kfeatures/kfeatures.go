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

		var chValue []string
		chValue = []string{chAttribute + " " + value}	

		if !kittyc.ModifyingAtLine(chAttribute, chValue[0]) {
			kittyc.WritingAtLine(section, chValue)
		}

		return nil
	}

	return empty
}

func ChangingMultipleValues(attribute string, values []string, section string) error {
	if attribute != "" && values != nil && section != "" {
		chAttribute := strings.ReplaceAll(attribute, " ", "_")

		if !kittyc.ModifyMultipleLines(chAttribute, values) {
			kittyc.WritingAtLine(section, values)
		}

		return nil
	}

	return empty
}
