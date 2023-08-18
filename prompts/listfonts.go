package prompts

import (
	"bufio"	
	"log"
	"os"
	"os/exec"
	"strings"
)

// listAllFontStyles: list specified font styles according to a selected font.
func listAllFontStyles(font string) []string {

	cmdFont := "fc-list : family style spacing outline scalable | grep -e spacing=100 -e spacing=90 | grep -e outline=True | grep -e scalable=True | grep '" + font + "'"

	cmd := exec.Command("bash", "-c", cmdFont)

	output, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.ReplaceAll(string(output), ":spacing=100:outline=True:scalable=True", " ")
	modifiedLines := strings.ReplaceAll(lines, ":style=", " ")
	fontStyles := strings.Split(modifiedLines, "\n")	

	return fontStyles[:len(fontStyles)-1]
}

// identifyFont: reads the kitty.conf file for getting the font_family value and saves its styles inside of a list.
func identifyFont() []string {
	homePath, _ := os.UserHomeDir()
	kittyConfPath := homePath + "/.config/kitty/kitty.conf"

	file, err := os.OpenFile(kittyConfPath, os.O_RDONLY, 0400)
	if err != nil {
		log.Fatalln(err)
		return []string{}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var identifiedFont []string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "font_family") {
			font := strings.Split(line, " ")

			styles := listAllFontStyles(font[1])	

			identifiedFont = append(identifiedFont, styles...)
			return identifiedFont
		}
	}
	
	return []string{""} 
}

// specificFontStyles: gets a new list based on the identified font styles and the current value from answers.Option (selected "change attribute"). 
func specificFontStyles(input string, list []string) []string {
	var adaptedStyles []string

	tempInput := strings.ReplaceAll(input, " font", "")

	for _, value := range list {
		tempValue := strings.ToLower(value)
		if strings.Contains(tempValue, tempInput) {
			adaptedStyles = append(adaptedStyles, "auto", strings.Title(tempValue), "none")
			return adaptedStyles
		}
	}

	return []string{"auto", "none"}
}
