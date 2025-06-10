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

// возьмём код из примера 2 и немного его изменим добавив цвета

func main() {
	// =========================================================================== loading

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

	// ======================================================================= create font

	newFontName := "new mixed font"
	fontMix := lascii.FontMixSettings_t{
		NewName:    newFontName,
		BaseName:   "style_2_letters",
		ShadowName: "style_2_shadow_2",
		Offset:     lascii.Vec2_t{X: -6, Y: -3},
	}

	err = lascii.CreateFontMix(fontMix)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	// ======================================================================= setting SGR

	// сделаем нашу собственную функцию
	// которая будет выдавать различные варианты раскраски
	lsb := createANSIcodes()

	// ============================================================================= print

	// чтобы вывести все варианты немного поменяем createArt добваив новый аргумент
	// и будем вызывать всё это в цикле
	for i := range lsb {
		fmt.Printf("================================================================================== variant %d\n", i)
		printArt(createArt(newFontName, lsb[i]))
	}
	fmt.Printf("==================================================================================\n")

}

// ___________________________________________________________________________SUBFUNCTIONS

func createANSIcodes() []lascii.LSB_t {
	lsb := make([]lascii.LSB_t, 0, 10)
	newLSB := lascii.LSB_t{}

	// для начала
	// разукрашивать мы будем используюя ANSI последовательности,
	// а точнее только SGR таблицу (смотри ссылку и ищи эту таблицу)
	// https://ru.wikipedia.org/wiki/%D0%A3%D0%BF%D1%80%D0%B0%D0%B2%D0%BB%D1%8F%D1%8E%D1%89%D0%B8%D0%B5_%D0%BF%D0%BE%D1%81%D0%BB%D0%B5%D0%B4%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE%D1%81%D1%82%D0%B8_ANSI

	// в lascii есть две функции для работы с этой таблицой
	// SetSGRParam и ConvertColor2SGRValue
	// они позволяют настроить всю талицу для вывода

	// что значит lsb?
	// letter - сами буквы
	// shadow - тень
	// back - фон у букв
	//
	// три слоя на которые длится шрифт
	// в примере 2 уже далали shadow, а back есть всегда
	// в примерах ниже мы их увидим

	// теперь перейдём к настройкам

	// variant 0 -------------------------------------------------------------------------
	// подробнее разберём что происходит

	// разукрасим слой letter в цвет CYAN
	//                  layer             param             value
	lascii.SetSGRParam(&newLSB.L, lascii.SGR_FONT_CYAN, lascii.SGR_PARAM_ON)

	// а слой shadow покрасим в rgb(0, 100, 100);
	// данный цвет не зарезервирован в SGR таблице
	// поэтому чтобы его поставить нам понадобится ConvertColor2SGRValue

	// выбрали цвет
	color := lascii.Color_t{R: 0, G: 80, B: 80}
	//                   цветовые моды описаны в SGR таблице код 38 или 48
	// если мод SGR_COLOR_MODE_STANDART то он находится в параметре R
	val := lascii.ConvertColor2SGRValue(lascii.SGR_COLOR_MODE_EXTENDED, color)
	// после настройки цвета нужно его передать в качестве значения
	lascii.SetSGRParam(&newLSB.S, lascii.SGR_FONT_COLOR, val)
	//                         |
	//               применяем цвет к shadow

	// включим UNDERLINE для символов из которых состоит слой shadow
	lascii.SetSGRParam(&newLSB.S, lascii.SGR_UNDERLINE, lascii.SGR_PARAM_ON)

	lsb = append(lsb, newLSB)
	newLSB = lascii.LSB_t{} // обнуляем

	// variant 1 -------------------------------------------------------------------------

	// letter rgb(0, 100, 250);
	// shadow rgb(40, 250, 0);

	color = lascii.Color_t{R: 0, G: 100, B: 250}
	val = lascii.ConvertColor2SGRValue(lascii.SGR_COLOR_MODE_EXTENDED, color)
	lascii.SetSGRParam(&newLSB.L, lascii.SGR_FONT_COLOR, val)

	color = lascii.Color_t{R: 40, G: 250, B: 0}
	val = lascii.ConvertColor2SGRValue(lascii.SGR_COLOR_MODE_EXTENDED, color)
	lascii.SetSGRParam(&newLSB.S, lascii.SGR_FONT_COLOR, val)

	lsb = append(lsb, newLSB)
	// не буду обнулять

	// variant 2 -------------------------------------------------------------------------

	// просто поменяю цвет тени

	// letter rgb(0, 100, 250);
	// shadow rgb(255, 0, 128);

	color = lascii.Color_t{R: 255, G: 0, B: 128}
	val = lascii.ConvertColor2SGRValue(lascii.SGR_COLOR_MODE_EXTENDED, color)
	lascii.SetSGRParam(&newLSB.S, lascii.SGR_FONT_COLOR, val)

	lsb = append(lsb, newLSB)
	// не буду обнулять

	// variant 3 -------------------------------------------------------------------------

	// выключу цвет для слоя letter

	// letter rgb(255, 255, 255);
	// shadow rgb(255, 0, 128);

	lascii.SetSGRParam(&newLSB.L, lascii.SGR_FONT_COLOR, lascii.SGR_PARAM_OFF)

	lsb = append(lsb, newLSB)
	// не буду обнулять

	// variant 4 -------------------------------------------------------------------------

	// поменяю цвет фона для back

	// letter rgb(255, 255, 255);
	// shadow rgb(255, 0, 128);
	// back rgb(50, 0, 128);

	color = lascii.Color_t{R: 50, G: 0, B: 128}
	val = lascii.ConvertColor2SGRValue(lascii.SGR_COLOR_MODE_EXTENDED, color)
	lascii.SetSGRParam(&newLSB.B, lascii.SGR_BACK_COLOR, val)

	lsb = append(lsb, newLSB)

	return lsb
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func createArt(fontName string, lsb lascii.LSB_t) [][]rune {
	art, err := lascii.CreateArt([]rune("LASCII"), fontName, &lsb)
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
