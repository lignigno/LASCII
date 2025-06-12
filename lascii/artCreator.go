package lascii

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// _______________________________________________________________________________SUBFUNCS

func getGroupSetting(zone1, zone2 SGRValue_t, unset SGRParam_t) string {
	groupSetting := ""

	if uint32(zone1) != 0 && uint32(zone2) == 0 {
		groupSetting += strconv.Itoa(int(unset))
	} else if zone1 != zone2 {
		if (unset == _SGR_UNSET_GROUP_FONT_COLOR ||
			unset == _SGR_UNSET_GROUP_BACK_COLOR) &&
			0 < (zone2&0xff) && (zone2&0xff) <= SGRValue_t(SGR_COLOR_MODE_STANDART) {
			if unset == _SGR_UNSET_GROUP_FONT_COLOR {
				groupSetting += "38;" + strconv.Itoa(int(zone2&0xff)) + ";"
			} else {
				groupSetting += "48;" + strconv.Itoa(int(zone2&0xff)) + ";"
			}
			groupSetting += strconv.Itoa(int(zone2 >> 8 & 0xff))

			if (zone2 & 0xff) == SGRValue_t(SGR_COLOR_MODE_EXTENDED) {
				groupSetting += ";" + strconv.Itoa(int(zone2>>16&0xff))
				groupSetting += ";" + strconv.Itoa(int(zone2>>24&0xff))
			}
		} else {
			groupSetting += strconv.Itoa(int(zone2))
		}
	}

	return groupSetting
}

// --------------------------------------------------------------------------------------|

func getChange(SettingLeft, SettingRight SGRSettings_t) string {
	change := "\033["

	for groupID := range _NUM_SGR_GROUPS {
		zone1 := SettingLeft[groupID]
		zone2 := SettingRight[groupID]
		groupSetting := getGroupSetting(zone1, zone2, _sgrGroups[groupID].Unset)

		if len(groupSetting) > 0 {
			if len(change) > 2 {
				change += ";"
			}
			change += groupSetting
		}
	}

	return change + "m"
}

// --------------------------------------------------------------------------------------|

func createLSBChanges(lsb *LSB_t) (_LSBTabe_t, bool) {
	empty := SGRSettings_t{}
	changesExitst := true

	// changeID = 4 * prev + current
	//  0  1  2  3    4  5  6  7    8  9 10 11   12 13 14 15
	// 00 0l 0s 0b | l0 ll ls lb | s0 sl ss sb | b0 bl bs bb
	//  0  l  s  b    3  0 ls lb    3 sl  0 sb    3 bl bs  0

	changes := _LSBTabe_t{
		//0                   1                        2                        3
		"", getChange(empty, lsb.L), getChange(empty, lsb.S), getChange(empty, lsb.B),
		//4                   5                        6                        7
		_SGR_RESET_SEQUENCE, "", getChange(lsb.L, lsb.S), getChange(lsb.L, lsb.B),
		//8                   9                       10                       11
		_SGR_RESET_SEQUENCE, getChange(lsb.S, lsb.L), "", getChange(lsb.S, lsb.B),
		//12                 13                       14                       15
		_SGR_RESET_SEQUENCE, getChange(lsb.B, lsb.L), getChange(lsb.B, lsb.S), "",
	}

	count := 0
	for i := range changes {
		if len(changes[i]) > 3 {
			count++
		}
	}
	if count == 0 {
		changes = _LSBTabe_t{}
		changesExitst = false
	}

	return changes, changesExitst
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func applyAlign(text []rune, pos []Vec2_t, size *_Size2_t, lu Vec2_t) {
	if size.W > _printSettings.MaxWidth {
		return
	}

	free := _printSettings.MaxWidth - size.W
	offset := Vec2_t{}
	numSpaces := 0

	for i := 0; _printSettings.Align == ALIGN_JUSTIFY && i < len(text); i++ {
		if text[i] == ' ' {
			numSpaces++
		}
	}

	if _printSettings.Align == ALIGN_CENTER {
		offset.X = free / 2
	} else if _printSettings.Align == ALIGN_RIGHT {
		offset.X = free
	}

	for i := range text {
		pos[i].X += offset.X - lu.X
		pos[i].Y += offset.Y - lu.Y

		if _printSettings.Align == ALIGN_JUSTIFY && text[i] == ' ' {
			offset.X += free / numSpaces
			free -= free / numSpaces
			numSpaces--

			continue
		}
	}

	if _printSettings.Align == ALIGN_RIGHT || _printSettings.Align == ALIGN_JUSTIFY {
		size.W = _printSettings.MaxWidth
	} else if _printSettings.Align == ALIGN_CENTER {
		size.W += offset.X
	}
}

// --------------------------------------------------------------------------------------|

func calcMarkup(font *_Font_t, text []rune) ([]rune, _Size2_t, []Vec2_t) {
	mayBePrinting := make([]rune, 0, len(text))
	canvas := _Size2_t{}
	positions := make([]Vec2_t, 0, len([]rune(text))*2)
	cursor, pos, lu, rd := Vec2_t{}, Vec2_t{}, Vec2_t{}, Vec2_t{}

	for _, r := range text {
		char, ok := font.Runes[r]
		if !ok && font.NullRune != nil {
			char = *font.NullRune
		} else if !ok {
			continue
		}
		mayBePrinting = append(mayBePrinting, r)

		pos.X = cursor.X
		pos.Y = -(char.Letter.Size.H - 1)
		positions = append(positions, pos)

		if !(_printSettings.Align == ALIGN_JUSTIFY && r == ' ') {
			lu.Y = min(lu.Y, pos.Y)
			rd.X = max(rd.X, pos.X+char.Letter.Size.W-1)
		}

		if char.Shadow != nil {
			pos.X = cursor.X + font.Offset.X
			pos.Y = font.Offset.Y - (char.Shadow.Size.H - 1)

			if !(_printSettings.Align == ALIGN_JUSTIFY && r == ' ') {
				lu.X = min(lu.X, pos.X)
				lu.Y = min(lu.Y, pos.Y)
				rd.X = max(rd.X, pos.X+char.Shadow.Size.W-1)
				rd.Y = max(rd.Y, pos.Y+char.Shadow.Size.H-1)
			}
		}

		if !(_printSettings.Align == ALIGN_JUSTIFY && r == ' ') {
			cursor.X += char.Letter.Size.W + int(_printSettings.LetterSpace)
		}
	}

	canvas.W = rd.X - lu.X + 1
	canvas.H = rd.Y - lu.Y + 1

	applyAlign(mayBePrinting, positions, &canvas, lu)

	return mayBePrinting, canvas, positions
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func resizeWidth(char *_Rune_t, posCur *Vec2_t, posNext Vec2_t) _Rune_t {
	size := _Size2_t{}
	stretchedSpace := _Rune_t{}

	size.W = posNext.X - posCur.X - int(_printSettings.LetterSpace)
	if size.W < 0 {
		size.W = 0
	}

	size.H = char.Letter.Size.H
	stretchedSpace.Letter = resizeRune(char.Letter, size)
	if char.Shadow != nil {
		size.H = char.Letter.Size.H
		tmp := resizeRune(*(char.Shadow), size)
		stretchedSpace.Shadow = &tmp
	}

	return stretchedSpace
}

// --------------------------------------------------------------------------------------|

func fillShadows(font *_Font_t, text []rune, blueprint *_BluePrint_t, canvas [][]_CanvasCell) {
	var symbolRune rune
	var cell *_CanvasCell
	var luAngle Vec2_t

	for i := range text {
		char, ok := font.Runes[text[i]]
		// нету проверки на nil так как уже отфильтровалось в calcMarkup
		if !ok {
			char = *font.NullRune
		}

		if _printSettings.Align == ALIGN_JUSTIFY && text[i] == ' ' {
			char = resizeWidth(
				&char,
				&blueprint.LettersPos[i],
				blueprint.LettersPos[i+1])
		}

		if char.Shadow != nil {
			luAngle.X = blueprint.LettersPos[i].X + font.Offset.X
			luAngle.Y = blueprint.LettersPos[i].Y + (char.Letter.Size.H - 1)
			luAngle.Y += font.Offset.Y - (char.Shadow.Size.H - 1)

			for y := range char.Shadow.Size.H {
				for x := range char.Shadow.Size.W {
					cell = &canvas[luAngle.Y+y][luAngle.X+x]
					symbolRune = char.Shadow.Runes[y][x]

					if symbolRune != font.EmptyRuneS {
						cell.Det = _DET_S_RUNE
						cell.R = symbolRune
					} else if cell.Det != _DET_S_RUNE && cell.Det != _DET_L_RUNE {
						if font.EmptyRuneS != font.BackRune {
							cell.Det = _DET_S_EMPTY
							cell.R = symbolRune
						}
					}
				}
			}
		}
	}
}

// --------------------------------------------------------------------------------------|

func fillLetters(font *_Font_t, text []rune, blueprint *_BluePrint_t, canvas [][]_CanvasCell) {
	var symbolRune rune
	var cell *_CanvasCell
	var luAngle Vec2_t

	for i := range text {
		char, ok := font.Runes[text[i]]
		if !ok {
			char = *font.NullRune
		}

		if _printSettings.Align == ALIGN_JUSTIFY && text[i] == ' ' {
			char = resizeWidth(
				&char,
				&blueprint.LettersPos[i],
				blueprint.LettersPos[i+1])
		}

		luAngle.X = blueprint.LettersPos[i].X
		luAngle.Y = blueprint.LettersPos[i].Y

		for y := range char.Letter.Size.H {
			for x := range char.Letter.Size.W {
				cell = &canvas[luAngle.Y+y][luAngle.X+x]
				symbolRune = char.Letter.Runes[y][x]

				if symbolRune != font.EmptyRuneL {
					cell.Det = _DET_L_RUNE
					cell.R = symbolRune
				} else if cell.Det != _DET_S_RUNE && cell.Det != _DET_L_RUNE {
					if font.EmptyRuneL != font.BackRune {
						cell.Det = _DET_L_EMPTY
						cell.R = symbolRune
					}
				}
			}
		}
	}
}

// --------------------------------------------------------------------------------------|

func createRowCanvas(font *_Font_t, text []rune, blueprint *_BluePrint_t) {
	canvas := make([][]_CanvasCell, blueprint.Size.H)

	for y := range blueprint.Size.H {
		canvas[y] = make([]_CanvasCell, blueprint.Size.W)

		for x := range blueprint.Size.W {
			canvas[y][x].Det = _DET_BACK
			canvas[y][x].R = font.BackRune
		}
	}

	fillShadows(font, text, blueprint, canvas)
	fillLetters(font, text, blueprint, canvas)

	blueprint.RowCanvas = canvas
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

// BLOCK : createRowCanvas
//
// Used in MAIN FUNC \/
func calcRealWidths(blueprint *_BluePrint_t) {
	blueprint.RealWidths = make([]int, blueprint.Size.H)

	for y := range blueprint.Size.H {
		sum := 0
		if blueprint.ChangesExist {
			sum = len(_SGR_RESET_SEQUENCE)
		}
		prev := _DET_NULL
		current := _DET_NULL
		for x := range blueprint.Size.W {
			current = blueprint.RowCanvas[y][x].Det
			if current != prev {
				sum += len(blueprint.Changes[4*(prev/10)+(current/10)])
			}
			sum++

			prev = current
		}
		blueprint.RealWidths[y] = sum
	}
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func createCanvas(blueprint *_BluePrint_t) [][]rune {
	canvas := make([][]rune, blueprint.Size.H)

	for y := range blueprint.Size.H {
		canvas[y] = make([]rune, blueprint.RealWidths[y])

		prev := _DET_NULL
		current := _DET_NULL
		offset := 0
		for x := range blueprint.Size.W {
			current = blueprint.RowCanvas[y][x].Det

			if current != prev {
				changeID := 4*(prev/10) + (current / 10)
				curChange := blueprint.Changes[changeID]

				for _, r := range curChange {
					canvas[y][x+offset] = r
					offset++
				}
			}
			canvas[y][x+offset] = blueprint.RowCanvas[y][x].R

			prev = current
		}
		if blueprint.ChangesExist {
			for _, r := range _SGR_RESET_SEQUENCE {
				canvas[y][blueprint.Size.W+offset] = r
				offset++
			}
		}
	}

	return canvas
}

// ______________________________________________________________________________MAIN FUNC

func CreateArt(text []rune, fontName string, lsb *LSB_t) ([][]rune, error) {
	font, ok := _fontsLib.Fonts[fontName]
	if !ok {
		return nil, errors.New("font {" + fontName + "} not exist")
	}
	if lsb == nil {
		return nil, errors.New("lsb is nil")
	}

	if _printSettings.Align == ALIGN_JUSTIFY {
		text = []rune(strings.Trim(string(text), " "))
		re := regexp.MustCompile(`[ ]+`)
		text = []rune(re.ReplaceAllString(string(text), " "))
	}

	blueprint := _BluePrint_t{}
	blueprint.Changes, blueprint.ChangesExist = createLSBChanges(lsb)
	text, blueprint.Size, blueprint.LettersPos = calcMarkup(&font, text)

	if blueprint.Size.W > _printSettings.MaxWidth {
		textSize := strconv.Itoa(blueprint.Size.W)
		maxWidth := strconv.Itoa(_printSettings.MaxWidth)
		return nil, errors.New("overflow terminal (" + textSize + " > " + maxWidth + ")")
	}

	createRowCanvas(&font, text, &blueprint)
	calcRealWidths(&blueprint)

	canvas := createCanvas(&blueprint)

	return canvas, nil
}
