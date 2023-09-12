package kfeatures

import (
	"os"
	"testing"
)

func TestReplacingKittyFile(t *testing.T) {
	homeFilePath, _ := os.UserHomeDir()

	filePath1 := homeFilePath + "/Downloads/kitty_extensive.conf"

	if fileTheme, err := ReplacingKittyFile(filePath1); !fileTheme {
		t.Errorf("Kitty theme is not replaced in kitty.conf file: %t, %s", fileTheme, err)
	} 
}
