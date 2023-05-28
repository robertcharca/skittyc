package kfeatures

import (
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
)

func AutomaticNewFont (font string) {
	// Download a font or search if this font already exists
}

func SetNewFont (font string) {
	fontAttribute := "font_family"

	fontValue := fontAttribute + " " + font
	
	if !kittyc.ModifyingAtLine(fontAttribute, fontValue) {
		kittyc.WritingAtLine("# Fonts", fontValue)
	} 
}

func ChangingFontValues (attribute, value string) {	
	fontAttribute := strings.ReplaceAll(attribute, " ", "_")

	var fontValue string
	fontValue = fontAttribute + " " + value	

	if !kittyc.ModifyingAtLine(fontAttribute, fontValue) {
		kittyc.WritingAtLine("# Fonts", fontValue)
	}
}

