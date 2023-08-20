package prompts

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

var cursorSetQuestion = []*survey.Question{
	{
		Name: "setCursor",
		Prompt: &survey.Select{
			Message: "What would you like to set in your cursor?",
			Options: []string{"cursor", "cursor shape", "cursor text color"},
		},
	},
}

var cursorChangeQuestion = []*survey.Question{
	{
		Name: "changeCursor",
		Prompt: &survey.Select{
			Message: "Choose an option: ",
			Options: []string{"cursor beam thickness", "cursor underline thickness", "cursor blink interval", 
				"cursor stop blinking after"},
		},
	},
}

func HandleSetCursor() (string, string) {
	answers := struct{ Option string `survey:"setCursor"`}{}
	
	err := survey.Ask(cursorSetQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var existingCursor string

	cursorSetColor := &survey.Input{
		Message: answers.Option,
	}

	cursorSelectType := &survey.Select{
		Message: "Select your cursor shape: ",
		Options: []string{"block", "beam", "underline"},
	}

	switch answers.Option {
	case "cursor shape":
		survey.AskOne(cursorSelectType, &existingCursor)
	default:
		survey.AskOne(cursorSetColor, &existingCursor, survey.WithValidator(hexCodeValidation.Validate))
	}	

	return answers.Option, existingCursor
}

func HandleChangeCursor() (string, string) {
	answers := struct{ Option string `survey:"changeCursor"`}{}
	
	err := survey.Ask(cursorChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var existingCursor string

	cursorInputNumber := &survey.Input{
		Message: answers.Option,
	}

	cursorInputNumberRange := &survey.Input{
		Message: answers.Option,
	}

	switch answers.Option{
	case "cursor blink interval":
		survey.AskOne(cursorInputNumberRange, &existingCursor, survey.WithValidator(numberAllRanges.Validate))
	default:
		survey.AskOne(cursorInputNumber, &existingCursor, survey.WithValidator(numberPositiveOnly.Validate))
	}

	return answers.Option, existingCursor
}
