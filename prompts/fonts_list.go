package prompts

import (
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

var answers = struct{ Option string `survey:"changeFont"`}{}

func HandleFontChangeValues () (string, string) {
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
