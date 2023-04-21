package kittyc

type ThemesInformation struct {
	project string
	name string
	license string
	link string
	description string
}

type Background struct {
	background string
	foreground string
	selection_background string
	selection_foreground string
}

type Color struct {
	colorn string
}

type KittyConfStructure struct {
	author ThemesInformation
	
	background Background

	color Color
}


