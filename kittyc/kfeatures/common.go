package kfeatures

import (
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
)

func ChangingValues(attribute, value , section string) {
	chAttribute := strings.ReplaceAll(attribute, " ", "_")

	var chValue string
	chValue = chAttribute + " " + value	

	if !kittyc.ModifyingAtLine(chAttribute, chValue) {
		kittyc.WritingAtLine(section, chValue)
	}
}
