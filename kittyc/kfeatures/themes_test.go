package kfeatures

import (	
	"testing"
)

type kittyThemeFile struct {
	filePath string
	fileName string
}

func TestReplacingKittyFile(t *testing.T) {	
	filePath1 := homePath + "/Downloads/tokyo-night-kitty.conf"

	if err := ReplacingKittyFile(filePath1); err != nil {
		t.Errorf("Kitty theme is not replaced in kitty.conf file: %s", err)
	} 
}
