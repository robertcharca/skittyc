package kittyc

import (
	"log"
	"os/exec"
	"strings"
)

// ListAllFonts: list all fonts through a command that gets all monospace fonts.
func ListAllFonts() []string{	
	cmd := exec.Command("bash", "-c", "fc-list : family spacing outline scalable | grep -e spacing=100 -e spacing=90 | grep -e outline=True | grep -e scalable=True")

	output, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)	
	}
	
	lines := strings.ReplaceAll(string(output), ":spacing=100:outline=True:scalable=True", "")
	systemFonts := strings.Split(string(lines), "\n")

	for i := 0; i < len(systemFonts); i++ {
		if strings.Contains(systemFonts[i], ",") {
			tempList := strings.Split(systemFonts[i], ",")
			tempResult := strings.TrimSpace(tempList[0])
			systemFonts[i] = tempResult
		}
	}

	return systemFonts[:len(systemFonts)-1]	
}
