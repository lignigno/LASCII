package main

import (
	"fmt"
	"os"

	"github.com/lignigno/LASCII/v2/lascii"
)

const (
	FONTS_DIR = "../../fonts"
)

// ______________________________________________________________________________MAIN FUNC

func main() {

	// загружаю шрифты которые буду миксовать
	loadSettings := lascii.LoadingSettings_t{
		Path:      FONTS_DIR + "/" + "super_fonts/style_2",
		Width:     0,
		Height:    0,
		HardASCII: false,
	}

	err := lascii.LoadFonts(loadSettings)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("==========\n")
	checkFonts()
	fmt.Printf("\nBefore /\\\n")
	fmt.Printf("==========\n")

	// чтобы создать новый шрифт с тенью укажем какие шрифты будем использовать
	newFontName := "new mixed font"
	fontMix := lascii.FontMixSettings_t{
		NewName:    newFontName,       // как назовём новый шрифт
		BaseName:   "style_2_letters", // основной шрифт, все настройки берутся из него
		WithShadow: true,
		ShadowName: "style_2_shadow_2",          // какой шрифт будет тенью
		Offset:     lascii.Vec2_t{X: -6, Y: -3}, // на сколько тень будет смещаться
	}

	// уточню на счёт смешения
	// смещение будет происходить для каждой буквы по отдельности
	// тень будет смещатся от нижнего левого угла

	//
	//            font base                shadow font
	//
	//         o--------------o
	//         |              |
	//         |              |
	//         |              |
	//         |              |
	//         |              |         ################
	//         |              |         ##            ##
	//         |              |         ##            ##
	//         o--------------o         ################
	//                     \               /
	//                      \             /
	//                       \  offset   /
	//                        \  (8;4)  /
	//                         \       /
	//                          \     /
	//                           \   /
	//                            mix
	//
	//                    o--------------o
	//                    |       #######|########
	//                    |       ##     |      ##
	//                    |       ##     |      ##
	//  4y offset  --  >  |       #######|########
	//             *      |              |
	//             *      |              |
	//             --  >  |              |
	//                    o--------------o
	//                     |      |
	//                     |******|  8x offset

	err = lascii.CreateFontMix(fontMix)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("==========\n")
	fmt.Printf("After  \\/\n\n")
	// обратите внимание что появился новый шрифт
	checkFonts()
	fmt.Printf("==========\n")

	// распечатаем "LASCII" нашим новым созданым шрифтом
	// использую для этого функции созданые в примере 0
	printArt(createArt(newFontName))
}

// ___________________________________________________________________________SUBFUNCTIONS

func checkFonts() {
	fontNames := lascii.GetFontNames()

	for i, name := range fontNames {
		fmt.Printf("font[%2d] {%s}\n", i, name)
	}
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func createArt(fontName string) [][]rune {
	art, err := lascii.CreateArt([]rune("LASCII"), fontName, &lascii.LSB_t{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	return art
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func printArt(art [][]rune) {
	for i := 0; i < len(art); i++ {
		fmt.Printf("|%s|\n", string(art[i]))
	}
}
