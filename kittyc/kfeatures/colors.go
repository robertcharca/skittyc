package kfeatures

import "github.com/robertcharca/skittyc/kittyc"

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

func SetColors(path string) bool {
	return false
}
