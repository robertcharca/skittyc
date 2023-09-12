package kfeatures

import (
	"bufio"	
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
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
	{"# Other", []string{"wayland_", "macos_", "map", "kitty_", "scrollback", "include", "shell", "listen", "allow_remote_control"}},
}

var homePath, _ = os.UserHomeDir()
var kittyPath = homePath + "/.config/kitty/kitty.conf"
var kittyProfile = homePath + "/.config/kitty/profile/"

func filteringKittyTheme(path string) ([]string, error) {	
	sectionLines := make(map[string][]string)
	var completeLines []string
	var noEmptyList []string
	var currentSection string

	// Getting lines that are not comments from the file
	file, err := os.OpenFile(path, os.O_RDWR, 0644)			
	if err != nil {	
		return []string{}, err
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

	specificSectionOrder := kittyc.FileStructureTitle

	// Adding missing sections in the map
	for _, section := range specificSectionOrder {
		_, ok := sectionLines[section] 

		if !ok {
			sectionLines[section] = append(sectionLines[section], "")
		}	
	}

	// Transfering the map values into a slice according to the file structure
	for i := 0; i < len(specificSectionOrder); i++ {
		for section, lines := range sectionLines {
			if specificSectionOrder[i] == section {
				completeLines = append(completeLines, section)
				completeLines = append(completeLines, lines...)
			}
		}
	}

	// Removing empty strings
	for i := 0; i < len(completeLines); i++ {
		emptyLine := strings.TrimSpace(completeLines[i])

		if emptyLine != "" {
			noEmptyList = append(noEmptyList, completeLines[i])
		}
	}	

	return noEmptyList, nil 
}

func rewrittingKittyFile(path string, newValues []string) error {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {	
		return err
	}

	truncate := file.Truncate(0) 
	if truncate != nil {	
		return truncate
	}

	_, offsErr := file.Seek(0, io.SeekStart)
	if offsErr != nil {	
		return offsErr
	}
	
	writer := bufio.NewWriter(file)	

	for i := 0; i < len(newValues); i++ {
		section := strings.HasPrefix(newValues[i], "#")

		if section && newValues[i] != "# Fonts" {
			if _, err := writer.WriteString("\n" + newValues[i] + "\n"); err != nil {
				return err
			}
		} else {
			if _, err := writer.WriteString(newValues[i] + "\n"); err != nil {
				return err
			}
		} 
	}

	writer.Flush()
	
	return nil 
}

func ReplacingKittyFile(path string) error {	
	kittyThemeRepl, err := filteringKittyTheme(path)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if err := rewrittingKittyFile(kittyPath, kittyThemeRepl); err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
