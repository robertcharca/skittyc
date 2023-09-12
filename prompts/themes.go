package prompts

import "github.com/AlecAivazis/survey/v2"

func ProfileNameInput() string {
	newFileName := inputSurvey("Type your profile name: ", survey.MinLength(1))

	return newFileName
}
