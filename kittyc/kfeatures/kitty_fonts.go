package kfeatures

import (
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
)

func SetNewFont () {
	//
}

func ChangingFontValues (attribute, value string) {	
	fontAttribute := strings.ReplaceAll(attribute, " ", "_")

	var fontValue string
	fontValue = fontAttribute + " " + value	

	if !kittyc.ModifyingAtLine(fontAttribute, fontValue) {
		kittyc.WritingAtLine("# Fonts", fontValue)
	}
}

