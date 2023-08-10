package kfeatures

import (
	"fmt"	
	"strings"

	"github.com/robertcharca/skittyc/kittyc"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// verifyFontDownload: compares three download alternatives and checks if the status is between 200 and 299
func verifyFontDownload(font string) kittyc.UrlDownload {
	corrFont := strings.ReplaceAll(font, " ", "-")

	firstUrl := "https://www.1001fonts.com/download/" + corrFont + ".zip"
	secondUrl := "https://www.fontsquirrel.com/fonts/download/" + corrFont
	
	urlAlternatives := []kittyc.UrlDownload{
		{
			Link: firstUrl,
			Format: ".zip",
			DownloadPath: "/.local/share/fonts/",
		},
		{
			Link: secondUrl,
			Format: ".zip",
			DownloadPath: "/.local/share/fonts/",
		},
		{
			Link: font,
			Format: ".zip",
			DownloadPath: "/.local/share/fonts/",
		},
	}

	for _, alts := range urlAlternatives {
		_, status := alts.VerifyDownload()

		if status {
			return alts
		}
	}	

	return kittyc.UrlDownload{}
}

func DownloadNewFont(font string) string {
	var fontName string

	fontDownloadUrl := verifyFontDownload(font)	
	file, downloaded, path := kittyc.DownloadFile(fontDownloadUrl) 

	if downloaded {
		kittyc.UnzipFile(path, file)
		fmt.Println("Unzipped file. Check it out!")
		fontName = strings.ReplaceAll(file, ".zip", "")
		newFN := strings.ReplaceAll(fontName, "-", " ")
		fmt.Printf("newFN: %s\n", newFN)
		return newFN 
	} else {
		fmt.Println("Problem. Check it out!")
	}

	return ""
}

func SetFontComparing(font string) {
	var lowerFonts []string

	entryFont := strings.ToLower(font)
	editedFonts := kittyc.ListAllFonts()	

	for _, v := range(editedFonts) {
		lower := strings.ToLower(v)
		lowerFonts = append(lowerFonts, lower)
	}	

	fonts, _ := kittyc.SearchingSimilarValues(lowerFonts, entryFont)

	if !fonts {
		fmt.Println("Does this font exist?")
	} else {
		SetNewFont(cases.Title(language.English, cases.NoLower).String(entryFont))
		fmt.Println("Implemented font. Check it out")
	}
}

func SetNewFont(font string) {
	fontAttribute := "font_family"

	fontValue := fontAttribute + " " + font
	
	if !kittyc.ModifyingAtLine(fontAttribute, fontValue) {
		kittyc.WritingAtLine("# Fonts", fontValue)
	} 
}
