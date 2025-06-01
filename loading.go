package ponylascii

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// ___________________________________________________________________________SUBFUNCTIONS

func getDstsFromDir(path string) ([]_FileDestination_t, error) {
	if path[len([]rune(path))-1] != '/' {
		path += "/"
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	countFiles := 0
	for i := 0; i < len(entries); i++ {
		if entries[i].IsDir() {
			continue
		}

		countFiles++
	}

	res := make([]_FileDestination_t, countFiles)
	countFiles = 0
	for i := 0; i < len(entries); i++ {
		if entries[i].IsDir() {
			continue
		}

		tmpName := entries[i].Name()
		tmpExt := filepath.Ext(tmpName)
		res[countFiles].Dir = path
		res[countFiles].FileName = strings.TrimSuffix(tmpName, tmpExt)
		res[countFiles].Ext = tmpExt
		countFiles++
	}

	return res, nil
}

// --------------------------------------------------------------------------------------|

func getFileDsts(path string) ([]_FileDestination_t, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		res, err := getDstsFromDir(path)

		return res, err
	}

	res := make([]_FileDestination_t, 1)
	res[0].Dir = filepath.Dir(path)
	if len([]rune(res[0].Dir)) > 1 {
		res[0].Dir += "/"
	}
	res[0].Ext = filepath.Ext(path)
	res[0].FileName = strings.TrimSuffix(filepath.Base(path), res[0].Ext)

	return res, nil
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func removeExistingFonts(dsts []_FileDestination_t) []_FileDestination_t {

	countUnique := 0
	for i := 0; i < len(dsts); i++ {
		if _, ok := _fontsLib.Fonts[dsts[i].FileName]; !ok {
			countUnique++
		}
	}

	clearDsts := make([]_FileDestination_t, 0, countUnique)
	for i := 0; i < len(dsts); i++ {
		if _, ok := _fontsLib.Fonts[dsts[i].FileName]; !ok {
			clearDsts = append(clearDsts, dsts[i])
		}
	}

	return clearDsts
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func getRawFont(path string) ([]string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, _DEFAULT_PERMISSION)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	rawFont := []string(nil)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 {
				rawFont = append(rawFont, strings.TrimRight(line, "\r\n"))
			}
			break
		} else if err != nil {
			return rawFont, err
		}
		rawFont = append(rawFont, strings.TrimRight(line, "\r\n"))
	}

	return rawFont, nil
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func ConvertFontParams2Regexp(height int, HardASCII bool) string {
	rule := `^`

	if !HardASCII {
		rule += `(.)`
	}
	if height < 1 {
		if !HardASCII {
			rule += `[ ]*`
		}

		rule += `(\d{1,2})`
	}

	return rule + `$`
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func getCharParams(char rune, h int, HardASCII bool, params []string) (rune, int, error) {

	if !HardASCII {
		char = []rune(params[1])[0]
	}

	heightID := 0
	if h < 1 {
		heightID++

		if !HardASCII {
			heightID++
		}
	}

	if heightID > 0 {
		var err error
		h, err = strconv.Atoi(params[heightID])
		if err != nil {
			return char, h, errors.New("can't convert height {" + params[heightID] + "}")
		}
	}

	return char, h, nil
}

// --------------------------------------------------------------------------------------|

func parseChar(rawFont []string, i, width, height int) (_Rune_t, error) {
	newChar := _Rune_t{}

	baseChar := _BaseRune_t{}
	if i < len(rawFont) {
		if width < 1 {
			width = len([]rune(rawFont[i]))
		}

		baseChar.Runes = make([][]rune, height)
		for j := 0; j < height; j++ {
			baseChar.Runes[j] = make([]rune, 0, width)
		}
		baseChar.Runes[0] = append(baseChar.Runes[0], []rune(rawFont[i])...)

		i++
	}

	for y := 1; i < len(rawFont) && y < int(height); i, y = i+1, y+1 {
		row := []rune(rawFont[i])

		if len(row) != width {
			return newChar, errors.New("incorrect width")
		}

		baseChar.Runes[y] = append(baseChar.Runes[y], row...)
	}

	baseChar.Size.W = width
	baseChar.Size.H = height
	newChar.Letter = baseChar

	return newChar, nil
}

// --------------------------------------------------------------------------------------|

func parseFont(rawFont []string, width, height int, HardASCII bool) (_Font_t, error) {
	font := _Font_t{
		EmptyRuneL: ' ',
		EmptyRuneS: ' ',
		BackRune:   ' ',
		NullRune:   &_runeNull,
		Runes:      make(map[rune]_Rune_t),
	}

	rule := ConvertFontParams2Regexp(height, HardASCII)
	rgxpParam, err := regexp.Compile(rule)
	if err != nil {
		return font, err
	}

	heightSetting := height
	asciiCode := _FIRST_ASCII_CODE
	for i := 0; i < len(rawFont); i, asciiCode = i+1, asciiCode+1 {
		rawParams, charName := rgxpParam.FindStringSubmatch(rawFont[i]), rune(asciiCode)
		if rawParams == nil {
			return font, errors.New("incorrect char parameter {" + rawFont[i] + "}")
		}
		charName, height, err = getCharParams(charName, heightSetting, HardASCII, rawParams)
		if err != nil {
			return font, err
		}

		if _, ok := font.Runes[charName]; ok {
			return font, errors.New("rune {" + string(charName) + "} in font exist")
		}

		font.Runes[charName], err = parseChar(rawFont, i+1, width, height)
		if err != nil {
			return font, errors.New("{" + string(charName) + "} " + err.Error())
		}
		font.MaxSize.W = max(font.MaxSize.W, font.Runes[charName].Letter.Size.W)
		font.MaxSize.H = max(font.MaxSize.H, font.Runes[charName].Letter.Size.H)
		i += font.Runes[charName].Letter.Size.H
	}

	return font, nil
}

// ______________________________________________________________________________MAIN FUNC

func LoadFonts(fonts LoadingSettings_t) error {
	dsts, err := getFileDsts(fonts.Path)
	if err != nil {
		return err
	}
	dsts = removeExistingFonts(dsts)

	for _, dst := range dsts {
		path := dst.Dir + dst.FileName + dst.Ext
		rawFont, err := getRawFont(path)
		if err != nil {
			return err
		}

		newFont, err := parseFont(rawFont, fonts.Width, fonts.Height, fonts.HardASCII)
		if err != nil {
			return errors.New("{" + path + "} : " + err.Error())
		}

		_fontsLib.Fonts[dst.FileName] = newFont
	}

	return nil
}
