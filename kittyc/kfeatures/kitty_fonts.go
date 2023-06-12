package kfeatures

import (
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
)

func DownloadNewFont (font string) {
	//
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

