package kittyc

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var notFound = errors.New("Value not found in array")

func SearchingValue(list []string, value string) (bool, error) {
	for _, sv := range list {
		if sv == value {
			return true, nil
		} 
	}

	return false, notFound
}

func SearchingSimilarValues(list []string, svalue string) (bool, error) {	
	for l := 0; l <= len(list); l++ {
		if list[l][:4] == svalue[:4] {
			return true, nil
		}
	}
	return false, notFound
}

func ConvertStringToList(s string) []string {	
	var sList []string

	list := []rune(s)

	for _, v := range list { 
		sList = append(sList, string(v))
	}

	return sList
}

func GettingMultipleValues(path string, keyword string) ([]string, error) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)			
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)		
	
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, keyword) {
			lines = append(lines, line)
		}
	}

	scanningFile := scanner.Err()
	if scanningFile != nil {
		fmt.Println(scanningFile.Error())
		return []string{}, scanningFile
	}
	return lines, nil 
}

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
