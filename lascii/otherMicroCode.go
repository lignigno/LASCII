package ponylascii

// _________________________________________________________________________GET FONT NAMES

func GetFontNames() []string {
	fontNames := make([]string, 0, len(_fontsLib.Fonts))

	for key := range _fontsLib.Fonts {
		fontNames = append(fontNames, key)
	}

	return fontNames
}

// _____________________________________________________________________SET PRINT SETTINGS

func SetPrintSettings(settings PrintSettings_t) {
	_printSettings.Align = settings.Align
	_printSettings.LineSpace = settings.LineSpace
	_printSettings.LetterSpace = settings.LetterSpace
	_printSettings.MaxWidth = settings.MaxWidth
}
