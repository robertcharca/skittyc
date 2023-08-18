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

	if answers.Option != "cursor shape" {
		csInput := survey.AskOne(cursorSetColor, &existingCursor, survey.WithValidator(hexCodeValidation.Validate))
		if csInput != nil {
			log.Fatalln(csInput)
			return "", ""
		}

		return answers.Option, existingCursor
	}

	csInputSelect := survey.AskOne(cursorSelectType, &existingCursor)
	if csInputSelect != nil {
		log.Fatalln(csInputSelect)
		return "", ""
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

	if answers.Option != "cursor blink interval" {
		csInput := survey.AskOne(cursorInputNumber, &existingCursor, survey.WithValidator(numberPositiveOnly.Validate))
		if csInput != nil {
			log.Fatalln(csInput)
			return "", ""
		}

		return answers.Option, existingCursor
	}

	csInputRange := survey.AskOne(cursorInputNumberRange, &existingCursor, survey.WithValidator(numberAllRanges.Validate))
	if csInputRange != nil {
		log.Fatalln(csInputRange)
		return "", ""
	}

	return answers.Option, existingCursor
}
