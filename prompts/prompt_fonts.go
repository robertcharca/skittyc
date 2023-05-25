package prompts

import (
	"fmt"
	"io/ioutil"
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
			Options: []string{"automatic", "select from system"},	
		},
	},
}

func listAllFonts() []string{	
	var fontsContent []string

	fontsPath := "/usr/share/fonts/truetype/"
	
	currentDir, err := ioutil.ReadDir(fontsPath) 
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, value := range currentDir {
		fontsContent = append(fontsContent, value.Name())
	}

	return fontsContent[1:]
}

func HandleNewFont () (string, string) {
	answers := struct{ Option string `survey:"setNewFont"`}{}

	err := survey.Ask(fontSetNewQuestion, &answers)
	if err != nil {
		log.Fatalln(err)
	}

	var textInput string

	fontSetNewInput := &survey.Input{
		Message: "Write your desired font:",

	}

	var existingFont string

	listFonts := listAllFonts()
	fontSelectExisting := &survey.Select{
		Message: "Select an existing font:",
		Options: listFonts,
	}	

	switch answers.Option {
	case "automatic":
		survey.AskOne(fontSetNewInput, &textInput)
		return answers.Option, textInput // Recall why you needed answers.Option
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
