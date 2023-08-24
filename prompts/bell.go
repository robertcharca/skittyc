package prompts

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
)

var bellSetQuestion = []*survey.Question{
	{
		Name: "setBell",
		Prompt: &survey.Select{
			Message: "What would you like to set in your bell?",
			Options: []string{"enable audio bell", "visual bell color", "window alert on bell"},
		},
	},
}

var bellChangeQuestion = []*survey.Question{
	{
		Name: "changeBell",
		Prompt: &survey.Select{
			Message: "Choose an option: ",
			Options: []string{"bell path", "bell on tab", "visual bell duration"},
		},
	},
}

func HandleSetBell() (string, string) {
	answers := struct{ Option string `survey:"setBell"`}{}
	
	err := survey.Ask(bellSetQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}
	
	var bellValue string	

	switch answers.Option {
	case "visual bell color":
		bellValue = selectSurveyOptions("Select your option: ", []string{"hex color", "none"})	

		if bellValue != "none" {
			bellValue = inputSurvey(answers.Option, hexCodeValidation.Validate) 
		}	
	default:
		bellValue = selectSurveyOptions("Select your confirmation: ", []string{"yes", "no"})	
	}  
	
	return answers.Option, bellValue 
}

func HandleChangeBell() (string, string) {
	answers := struct{ Option string `survey:"changeBell"`}{}
	homePath, _ := os.UserHomeDir()

	err := survey.Ask(bellChangeQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}
	
	var bellValue string

	bellPathInput := &survey.Input{
		Message: "Paste a file path for your bell song:",
		Suggest: func(toComplete string) []string {
			var filteredFiles []string

			confFiles, _ := filepath.Glob(homePath + toComplete + "*.wav")
			textFiles, _ := filepath.Glob(homePath + toComplete + "*.oga")
			confFiles = append(confFiles, textFiles...)
			filteredFiles = append(filteredFiles, confFiles...)

			return filteredFiles
		},
	}	
	
	switch answers.Option{
	case "bell path":
		survey.AskOne(bellPathInput, &bellValue, survey.WithValidator(survey.MinLength(6))) 
	case "visual bell duration":	
		bellValue = inputSurvey(answers.Option, numberPositiveOnly.Validate) 
	default:
		bellValue = selectSurveyOptions("Select your confirmation: ", []string{"yes", "no", "custom"})	

		if bellValue == "custom" {
			bellValue = inputSurvey("Type a word or an emoji: ", survey.MinLength(1))
			
			if verify := regexp.MustCompile(`\s`).MatchString(bellValue); verify {
				return answers.Option, `"` + bellValue + `"`
			}
		}
	}

	return answers.Option, bellValue 
}
