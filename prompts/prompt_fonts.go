package prompts

import (
	"fmt"	
	"log"

	"github.com/AlecAivazis/survey/v2"	
)

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

var fontSetNewQuestion = []*survey.Question{
	{
		Name: "setNewFont",
		Prompt: &survey.Select{
			Message: "How would you like to set your new font?",
			Options: []string{"from name", "from url", "select from system"},	
		},
	},
}

func HandleNewFont () (string, string) {
	answers := struct{ Option string `survey:"setNewFont"`}{}

	err := survey.Ask(fontSetNewQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}	

	var existingFont string

	listFonts := listAllFonts()
	fontSelectExisting := &survey.Select{
		Message: "Select an existing font:",
		Options: listFonts,
	}	

	switch answers.Option {
	case "from name":
		//	
	case "from url":
		//
	case "select from system":
		survey.AskOne(fontSelectExisting, &existingFont)
		return answers.Option, existingFont
	default:
		fmt.Println("Select a valid option")
	}

	return answers.Option, ""	
}

func HandleFontChangeValues () (string, string) {
	answers := struct{ Option string `survey:"changeFont"`}{}

	err := survey.Ask(fontChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)	
	}

	var textInput string	

	fontChangeInput := &survey.Input{
		Message: answers.Option,
	}

	textInputErr := survey.AskOne(fontChangeInput, &textInput)

	if textInputErr != nil {
		log.Fatalln(textInput)
		return " ", " "
	}

	return answers.Option, textInput
}
