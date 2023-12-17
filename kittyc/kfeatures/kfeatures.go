package kfeatures

import (
	"errors"
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
)

var errEmpty = errors.New("empty values not allowed")

func ChangingValues(attribute, value, section string) error {
	if attribute != "" && value != "" && section != "" {
		chAttribute := strings.ReplaceAll(attribute, " ", "_")

		chValue := []string{chAttribute + " " + value}

		if !kittyc.ModifyingAtLine(chAttribute, chValue[0]) {
			kittyc.WritingAtLine(section, chValue)
		}

		return nil
	}

	return errEmpty
}

func ChangingMultipleValues(attribute, values []string, section string) error {
	if attribute != nil && values != nil && section != "" {
		var chAttribute []string

		for _, attr := range attribute {
			change := strings.ReplaceAll(attr, " ", "_")
			chAttribute = append(chAttribute, change)
		}

		if !kittyc.ModifyMultipleLines(chAttribute, values) {
			kittyc.WritingAtLine(section, values)
		}

		return nil
	}

	return errEmpty
}

func DownloadKittyFiles(link, fileFormat string) string {
	kittyDownload := kittyc.UrlDownload{
		Link:         link,
		Format:       fileFormat,
		DownloadPath: "/Downloads/",
	}

	file, downloaded, path := kittyc.DownloadFile(kittyDownload)

	if !downloaded {
		return ""
	}

	return path + file
}
