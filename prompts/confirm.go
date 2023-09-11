package prompts

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/robertcharca/skittyc/kittyc"
)

func ConfirmFontExistence(font string) (bool, bool) {
	fontExistence, _ := kittyc.SearchingValue(kittyc.ListAllFonts(), font)

	unknownFont := false

	fontConfirm := &survey.Confirm{
		Message: "This font doesn't exist. Do you want to install it?",
	}

	if fontExistence == false {
		survey.AskOne(fontConfirm, &unknownFont)
		return unknownFont, true
	}

	return unknownFont, false
}

func ConfirmKittyConfExistence() (string, bool) {
	_, kittyconf := kittyc.KittyConfigExistence()

	var kittyConfirmValue string

	kittyFileConfirm := &survey.Select{
		Message: "A kitty.conf file already exists in your system. What do you want to do?",
		Options: []string{"save it as a profile", "replace it", "none"},
	}

	if kittyconf {
		if err := survey.AskOne(kittyFileConfirm, &kittyConfirmValue); err != nil {
			log.Fatalln(err)
		}
	}

	return kittyConfirmValue, kittyconf
}
