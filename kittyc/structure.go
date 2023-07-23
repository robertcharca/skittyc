package kittyc

import (
	"fmt"
	"os"
)

var fileStructureTitle = []string{"# Fonts", "# Background and Foreground", "# Color"}

func DisplayStructure (file *os.File) {
	for _, values := range fileStructureTitle {
		_, err := fmt.Fprintf(file, "%s \n\n", values) 

		if err != nil {
			fmt.Println(err.Error())
			return 
		}
	}
}

