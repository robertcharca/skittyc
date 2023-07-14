package kfeatures

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func verifyAutomaticDownload (font string) (bool, string, bool) {
	var (
		corrFont string
		urlFirstAlternative string
		urlSecondAlternative string
	)

	corrFont = strings.ReplaceAll(font, " ", "-")

	urlFirstAlternative = "https://www.1001fonts.com/download/" + corrFont + ".zip"
	urlSecondAlternative = "https://www.fontsquirrel.com/fonts/download/" + corrFont

	respFirst, err := http.Get(urlFirstAlternative)
	if err != nil {
		log.Fatalln(err)
	}

	respSecond, err := http.Get(urlSecondAlternative)
	if err != nil {
		log.Fatalln(err)
	}

	if respFirst.StatusCode >= 200 && respFirst.StatusCode <= 299 {
		return true, urlFirstAlternative, true
	} else if respSecond.StatusCode >= 200 && respSecond.StatusCode <= 299 {
		return true, urlSecondAlternative, false
	}

	return false, "", false
}

func downloadFontZip (font string) (string, bool, string) {
	var (
		fileName string
		fontsPath string
	)

	homePath, _ := os.UserHomeDir()
	fontsPath = homePath + "/.local/share/fonts/"

	fontStatus, download, zip := verifyAutomaticDownload(font)

	if !fontStatus {
		fmt.Println("This font cannot be downloaded.")
		return "", false, ""
	}

	fileURL, err := url.Parse(download)
	if err != nil {
		log.Fatalln(err)
	}

	path := fileURL.Path
	segments := strings.Split(path, "/")

	if zip {
		fileName = segments[len(segments)-1]
	} else {
		fileName = segments[len(segments)-1] + ".zip"
	}
	
	file, err := os.Create(fontsPath + fileName)
	if err != nil {
		log.Fatalln(err)
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	resp, err := client.Get(download)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

	return fileName, true, fontsPath
}

func DownloadNewFont (font string) string {
	file, downloaded, path := downloadFontZip(font)

	if downloaded {
		kittyc.UnzipFile(file, path)
		fmt.Println("Unzipped file. Check it out!")
		return font
	} else {
		fmt.Println("Problem. Check it out!")
	}

	return ""
}

func SetFontComparing (font string) {
	var lowerFonts []string

	entryFont := strings.ToLower(font)
	editedFonts := kittyc.ListAllFonts()	

	for _, v := range(editedFonts) {
		lower := strings.ToLower(v)
		lowerFonts = append(lowerFonts, lower)
	}	

	fonts, _ := kittyc.SearchingValue(lowerFonts, entryFont)	

	if !fonts {
		fmt.Println("Does this font exist?")
	} else {
		SetNewFont(cases.Title(language.English, cases.NoLower).String(entryFont))
		fmt.Println("Implemented font. Check it out")
	}
}

func SetNewFont (font string) {
	fontAttribute := "font_family"

	fontValue := fontAttribute + " " + font
	
	if !kittyc.ModifyingAtLine(fontAttribute, fontValue) {
		kittyc.WritingAtLine("# Fonts", fontValue)
	} 
}

func ChangingFontValues (attribute, value string) {	
	fontAttribute := strings.ReplaceAll(attribute, " ", "_")

	var fontValue string
	fontValue = fontAttribute + " " + value	

	if !kittyc.ModifyingAtLine(fontAttribute, fontValue) {
		kittyc.WritingAtLine("# Fonts", fontValue)
	}
}

