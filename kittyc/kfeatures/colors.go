package kfeatures

import (
	"fmt"

	"github.com/robertcharca/skittyc/kittyc"
)

func DownloadColors(link string) string {
	colorDownload := kittyc.UrlDownload{
		Link: link,
		Format: ".conf",
		DownloadPath: "/Downloads/",
	}
	
	file, downloaded, path := kittyc.DownloadFile(colorDownload)

	if !downloaded {
		return ""
	}

	return path + file
}

func SetColors(path string) {
	colors, err := kittyc.GettingMultipleValues(path, "color")
	if err != nil {
		fmt.Println(err.Error())
	}	

	errValues := ChangingMultipleValues("color", colors, "# Colors")
	if errValues != nil {
		fmt.Println(err.Error())
	}
}
