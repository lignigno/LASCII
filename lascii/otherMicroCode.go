package lascii

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

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

// _________________________________________________________________________GET WIDTH TERM

func GetWidthTerm() (width int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(out), " ")
	width, err = strconv.Atoi(parts[1][:len(parts[1])-1])
	if err != nil {
		log.Fatal(err)
	}

	return
}
