package prompts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/robertcharca/skittyc/kittyc"
)

func ConfirmFontExistence (font string) (bool, bool) {
	fontExistence, _ := kittyc.SearchingValue(listAllFonts(), font)

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
