package lascii

import (
	"fmt"
	"regexp"
)

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func clearANSICode(art [][]rune) [][]rune {
	re := regexp.MustCompile(`\033\[.*?m`)

	for i, str := range art {
		art[i] = []rune(re.ReplaceAllString(string(str), ""))
	}

	return art
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func clearDetCanvas(canvas [][]_DetLayer) {
	for y := range canvas {
		for x := range canvas[y] {
			canvas[y][x] = _DET_NULL
		}
	}
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func createDetCanvas(art [][]rune) ([][]_DetLayer, int) {
	square := 0
	detCanvas := make([][]_DetLayer, len(art))
	for i := range detCanvas {
		detCanvas[i] = make([]_DetLayer, len(art[i]))
		square += len(art[i])
	}

	return detCanvas, square
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func findRune(art [][]rune, r *_Rune_t) ([]Vec2_t, []_Size2_t) {

	return nil, nil
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func fillDetCanvas(canvas [][]_DetLayer, poses []Vec2_t, sizes []_Size2_t) int {
	if poses == nil || sizes == nil {
		return 0
	}

	return 0
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func restoreText(art [][]rune, canvas [][]_DetLayer) string {
	return "ya sdelyal :)"
}

// ______________________________________________________________________________MAIN FUNC

func сonvertArt2text(art [][]rune) ([]Findings_t, int) {
	findings := make([]Findings_t, 0, 1)
	var filedSquare int

	art = clearANSICode(art)
	detCanvas, square := createDetCanvas(art)

	for fontName, font := range _fontsLib.Fonts {
		filedSquare = 0
		clearDetCanvas(detCanvas)

		for _, r := range font.Runes {
			poses, sizes := findRune(art, &r)
			filedSquare += fillDetCanvas(detCanvas, poses, sizes)
		}

		if filedSquare > 0 {
			match := Findings_t{
				fontName: fontName,
				text:     restoreText(art, detCanvas),
				coverage: filedSquare,
			}

			findings = append(findings, match)
		}
	}

	{ // some debug print
		img := ""
		for y := range detCanvas {
			for x := range detCanvas[y] {
				if detCanvas[y][x] != _DET_NULL {
					img += "\033[1;38;2;0;255;0;48;2;35;50;43m"
				}

				img += string(art[y][x])

				if detCanvas[y][x] != _DET_NULL {
					img += "\033[0m"
				}
			}
			img += string("\n")
		}
		fmt.Printf("%s", img)
	}

	return nil, square
}
