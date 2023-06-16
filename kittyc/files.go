package kittyc

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func UnzipFile(fileName string, path string) {
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
