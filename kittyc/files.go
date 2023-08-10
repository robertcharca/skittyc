package kittyc

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type UrlDownload struct {
	Link string
	Format string
	DownloadPath string
}

func (u UrlDownload) VerifyDownload() (int, bool) {
	resp, err := http.Get(u.Link)
	if err != nil {
		log.Fatalln(err)
	}
	
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return resp.StatusCode, true
	}

	return resp.StatusCode, false
}

func (u UrlDownload) VerifyFormat() (string, string, bool) {
	path := u.Link

	fileUrl, err := url.Parse(path)
	if err != nil {
		log.Fatalln(err)
	}

	urlPath := fileUrl.Path
	segments := strings.Split(urlPath, "/")

	lastPathVal := segments[len(segments)-1]

	if strings.Contains(lastPathVal, ".") {
		fileSep := strings.Split(lastPathVal, ".")

		return fileSep[0], "." + fileSep[1], true
	}

	return lastPathVal, "", false
}

func DownloadFile(urlPath UrlDownload) (string, bool, string) {
	var fileName string

	_, urlValidation := urlPath.VerifyDownload()
	homePath, _:= os.UserHomeDir()
	downloadsPath := homePath + urlPath.DownloadPath

	if !urlValidation {
		return "", false, ""
	}

	name, format, hasFileFormat := urlPath.VerifyFormat()

	if !hasFileFormat {
		fileName = name + urlPath.Format
	} else {
		fileName = name + format
	}

	fmt.Println("fileName: ", fileName)

	file, err := os.Create(downloadsPath + fileName)
	if err != nil {
		log.Fatalln(err)
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	resp, err := client.Get(urlPath.Link)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)
	return fileName, true, downloadsPath
}

func UnzipFile(path, fileName string) {
	dirName := strings.Split(fileName, ".zip")
	finalDirZip := dirName[0]

	zipPath := path + fileName
	zipFile, err := zip.OpenReader(zipPath)
	if err != nil {
		panic(err)
	}

	defer zipFile.Close()

	for _, file := range zipFile.File {	
		extractedPath := filepath.Join(
			path + finalDirZip,
			file.Name,
		)
		fmt.Println("Unzipping file: ", extractedPath)

		if file.FileInfo().IsDir() {
			fmt.Println("Directory created: ", extractedPath)
			if err := os.MkdirAll(extractedPath, os.ModePerm); err != nil {
				panic(err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(extractedPath), os.ModePerm); err != nil {
			panic(err)
		}

		fontZip, err := os.OpenFile(extractedPath, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, file.Mode())
		if err != nil {
			panic(err)
		}

		fontZipArchive, err := file.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(fontZip, fontZipArchive); err != nil {
			panic(err)
		}

		fontZip.Close()
		fontZipArchive.Close()
	}

	removeZipFile := os.Remove(path + fileName)
	if removeZipFile != nil {
		log.Println(removeZipFile)
	}

}
