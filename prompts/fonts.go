package prompts

import (	
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/robertcharca/skittyc/kittyc"
)

var fontSetNewQuestion = []*survey.Question{
	{
		Name: "setNewFont",
		Prompt: &survey.Select{
			Message: "How would you like to set your new font?",
			Options: []string{"automatic", "url", "select from system"},	
		},
	},
}

var fontChangeQuestion = []*survey.Question{
	{
		Name: "changeFont",
		Prompt: &survey.Select{
			Message: "Choose an option: ",
			Options: []string{"font size", "bold font", "italic font", "bold italic font"},
			Help: "helping",
		},
	},
}

// HandleNewFont: sets a "select" survey that list different options to set a new font. 
func HandleNewFont() (string, string) {
	answers := struct{ Option string `survey:"setNewFont"`}{}

	err := survey.Ask(fontSetNewQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}	

	var existingFont string

	listFonts := kittyc.ListAllFonts()
	fontSelectExisting := &survey.Select{
		Message: "Select an existing font:",
		Options: listFonts,
	}

	fontUrlDownload := &survey.Input{
		Message: "Paste a url for implementing your font:",
	}

	fontAutomatic := &survey.Input{
		Message: "Type a font:",
		Suggest: func(toComplete string) []string {
			fontList := kittyc.ListAllFonts()
			return fontList
		},
	}

	switch answers.Option {
	case "automatic":
		survey.AskOne(fontAutomatic, &existingFont)	
		return answers.Option, existingFont
	case "url":
		survey.AskOne(fontUrlDownload, &existingFont, survey.WithValidator(urlValidation.Validate))
		return answers.Option, existingFont
	case "select from system":
		survey.AskOne(fontSelectExisting, &existingFont)
		return answers.Option, existingFont
	default:
		fmt.Println("Select a valid option")
	}

	return answers.Option, ""	
}

// HandleFontChangeValues: sets a "select" survey for the font "change attributes".
// Based on the selected attribute, it will set another "select" survey for an attribute according to a implemented font. 
func HandleFontChangeValues() (string, string) {
	answers := struct{ Option string `survey:"changeFont"`}{}

	err := survey.Ask(fontChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)	
	}

	var fontStyle string
	
	// Survey for listing font styles according to a font	
	fontStyleList := identifyFont()
	specificInput := specificFontStyles(answers.Option, fontStyleList)
	
	fontChangeStyleList := &survey.Select{
		Message: answers.Option,
		Options: specificInput,
	}

	// Survey for font size input (numeric)
	fontSize := &survey.Input{
		Message: answers.Option,
	}

	// Changing survey inputs according to the "attribute"
	if answers.Option != "font size" {
		fsInput := survey.AskOne(fontChangeStyleList, &fontStyle)
		if fsInput != nil {
			log.Fatalln(fontStyle)
			return " ", " "
		}

		return answers.Option, fontStyle
	}

	fsInput := survey.AskOne(fontSize, &fontStyle)
	if fsInput != nil {
		log.Fatalln(fontStyle)
		return " ", " " 
	}

	return answers.Option, fontStyle
}
