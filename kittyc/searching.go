package kittyc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"	
)

func WritingAtLine(section string, addedLine string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)			
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)		
	
	var lines []string
	
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)

		if strings.Contains(line, section) {
			newLine := addedLine 
			lines = append(lines, newLine)
		}
	}

	scanningFile := scanner.Err()
	if scanningFile != nil {
		fmt.Println(scanningFile.Error())
	}

	/*
		Truncate: changes the size of the file. In this case, it change the last size to 0, 
		which means it's completely empty.
	*/
	truncate := file.Truncate(0) 
	if truncate != nil {
		fmt.Println(err.Error())
	}

	/*
		Seek: sets a new offset according to a "principal" offset. In this case, the new
		offset is 0 based on the "start" offset, which also is 0. This means the offset will
		be positioned at the complete beginning of the file.
	*/
	_, offsErr := file.Seek(0, io.SeekStart)
	if offsErr != nil {
		fmt.Println(err.Error())
	}
	
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
		fmt.Println(err.Error())
	}
	}
	
	flush := writer.Flush()
	if flush != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Everything's great")
}
