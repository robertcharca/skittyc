package kittyc

import (
	"testing"
)

func TestDownloadFile(t *testing.T) {
	fileDownload := []UrlDownload{
		{
			Link: "https://www.1001fonts.com/download/unispace.zip",
			DownloadPath: "/Downloads/",
		},
		{
			Link: "https://www.fontsquirrel.com/fonts/download/jetbrains-mono",
			Format: ".zip",
			DownloadPath: "/Downloads/",
		},
	}

	for _, file := range fileDownload {
		fileName, isDownloaded, filePath := DownloadFile(file)

		if !isDownloaded {
			t.Errorf("DownloadFile output - file: %s, downloaded: %t, format: %s\n", fileName, isDownloaded, filePath)
		}	
	}
}
