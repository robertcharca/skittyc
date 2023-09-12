package kfeatures

import "testing"

func TestSKFProf(t *testing.T) {
	kittyProfs := []kittyThemeFile{
		{homePath + "/Downloads/tokyo-night-kitty.conf", "tokyo-night"},
		{homePath + "/Downloads/frappe.conf", "frappe"},
		{homePath + "/Downloads/mocha.conf", ""},
		{homePath + "/Downloads/frappe.conf", "  "},
		{homePath + "/Downloads/mocha.conf", "mocha     "},
		{homePath + "/Downloads/atom-text.txt", "atom"},
		{homePath + "/Downloads/kitty-extensive.conf", "extensive-af"},
	}

	for _, file := range kittyProfs {
		if err := SavingKittyFileProfile(file.filePath, file.fileName); err != nil {
			t.Errorf("Kitty profile hasn't been implemented: %s", err)
		}
	}
}
