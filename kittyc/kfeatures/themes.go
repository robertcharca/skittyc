package kfeatures

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"	
)

type KittyStructure struct {
	Section string
	CommonPatterns []string
}

var kittyHints = []KittyStructure{
	{"# Fonts", []string{"font_", "_font"}},
	{"# Background and Foreground", []string{"background", "foreground", "_foreground", "_background", "background_", "foreground_"}},
	{"# Cursor", []string{"cursor", "cursor_", "cursor_text_color"}},
	{"# Colors", []string{`color\d+`, `mark\d+`}},
	{"# Mouse", []string{"mouse_", "_mouse", "url_", "_urls", "click_", "pointer_", "copy_on_select", "strip_trailing_spaces"}},
	{"# Tab bar", []string{"tab_", "_tab"}},
	{"# Terminal bell", []string{"bell_", "_bell"}},
	{"# Window layout", []string{"window", "window_", "_window", "_layouts", "_border", "_borders", "inactive_text_alpha", "_strategy"}},
	{"# Other", []string{"wayland_", "macos_"}},
}

func filteringKittyTheme(path string) []string {	
	sectionLines := make(map[string][]string)	
	var currentSection string
 
	// Getting lines that are not comments from the file
	file, err := os.OpenFile(path, os.O_RDWR, 0644)			
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)		

	for scanner.Scan() {
		line := scanner.Text()	

		hasComments := strings.HasPrefix(line, "#")

		// Filtering lines according to each section
		if !hasComments {
			for _, section := range kittyHints {
				for _, pattern := range section.CommonPatterns {
					regexComp := regexp.MustCompile(pattern)
					if regexComp.MatchString(line) {
						currentSection = section.Section
						break
					}
				}	
			}

			sectionLines[currentSection] = append(sectionLines[currentSection], line)
		}

	}	

	fmt.Println(sectionLines)

	return []string{} 
}

func ReplacingKittyFile(path string) {
	list := filteringKittyTheme(path)

	fmt.Println(list)
}
