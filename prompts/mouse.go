package prompts

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
)

var mouseSetQuestion = []*survey.Question{
	{
		Name: "setMouse",
		Prompt: &survey.Select{
			Message: "What would you like to set for your mouse?",
			Options: []string{"url color", "url style", "default pointer shape",
			"pointer shape when grabbed", "pointer shape when dragging", "select by word characters"},
		},
	},
}

var mouseChangeQuestion = []*survey.Question{
	{
		Name: "changeMouse",
		Prompt: &survey.Select{
			Message: "Choose an option: ",
			Options: []string{"detect urls", "click interval", "copy on select", 
				"mouse hide wait", "focus follows mouse", "strip trailing spaces"},
		},
	},
}

func HandleSetMouse() (string, string) {
	answers := struct{ Option string `survey:"setMouse"`}{}
	
	err := survey.Ask(mouseSetQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var mouseValue string

	switch answers.Option {
	case "url color":
		mouseValue = inputSurvey(answers.Option, hexCodeValidation.Validate)
	case "select by word characters":
		mouseValue = inputSurvey(answers.Option, survey.MinLength(1))
	case "default pointer shape", "pointer shape when grabbed", "pointer shape when dragging":
		mouseValue = selectSurveyOptions("Select your option: ", []string{"arrow", "beam", "hand"})	
	case "url style":
		mouseValue = selectSurveyOptions("Select your url style: ", []string{"none", "curly", "dashed", "dotted", "double", "straight"})
	}

	return answers.Option, mouseValue
}

func HandleChangeMouse() (string, string) {
	answers := struct{ Option string `survey:"changeMouse"`}{}
	
	err := survey.Ask(mouseChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var mouseValue string

	switch answers.Option {
	case "mouse hide wait", "click interval":
		mouseValue = inputSurvey(answers.Option, numberAllRanges.Validate)
	case "detect urls", "copy on select", "focus follows mouse":
		mouseValue = selectSurveyOptions("Select your option: ", []string{"yes", "no"})
	case "strip trailing spaces":
		mouseValue = selectSurveyOptions("Select your option: ", []string{"never", "smart", "always"})
	}

	return answers.Option, mouseValue
}
