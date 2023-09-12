package kittyc

import (
	"fmt"
	"os"
)

var FileStructureTitle = []string{"# Fonts", "# Background and Foreground", "# Cursor", 
	"# Colors", "# Mouse", "# Tab bar", "# Terminal bell", "# Window layout", "# Other"}

func DisplayStructure (file *os.File) {
	for _, values := range FileStructureTitle {
		_, err := fmt.Fprintf(file, "%s \n\n", values) 

		if err != nil {
			fmt.Println(err.Error())
			return 
		}
	}
}

