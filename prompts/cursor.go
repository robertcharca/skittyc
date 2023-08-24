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

	var cursorValue string	

	switch answers.Option {
	case "cursor shape":
		cursorValue = selectSurveyOptions("Select your cursor shape: ", []string{"block", "beam", "underline"})	
	default:
		cursorValue = inputSurvey(answers.Option, hexCodeValidation.Validate)	
	}	

	return answers.Option, cursorValue 
}

func HandleChangeCursor() (string, string) {
	answers := struct{ Option string `survey:"changeCursor"`}{}
	
	err := survey.Ask(cursorChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}
	
	var cursorValue string	

	switch answers.Option{
	case "cursor blink interval":
		cursorValue = inputSurvey(answers.Option, numberAllRanges.Validate)	
	default:	
		cursorValue = inputSurvey(answers.Option, numberPositiveOnly.Validate)
	}

	return answers.Option, cursorValue 
}
