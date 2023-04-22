package kittyc

import (
	"fmt"
	"os"
)

var fileStructureTitle = []string{"# Fonts", "# Background", "# Foreground", "# Color"}

type ThemesInformation struct {
	project string
	name string
	license string
	link string
	description string
}

type Background struct {
	background string
	foreground string
	selectionBackground string
	selectionForeground string
}

type Color struct {
	colorn map[string]string 
}

type Font struct {
	font_family string
	bold_font string
	italic_font string
	font_size float32

}

func displayStructure (file *os.File) {
	for _, values := range fileStructureTitle {
		_, err := fmt.Fprintf(file, "%s \n\n", values) 

		if err != nil {
			fmt.Println(err.Error())
			return 
		}
	}
}

