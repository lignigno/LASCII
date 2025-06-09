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
	// из примера 0 известно что нужно указать "настройки загрузки"
	// о них и о том что должно быть указано в файле шрифта будет вестись речь
	// в этом уроке мы не будем печатать сам текст
	// толко изучать как можно загрузить шрифты

	// ОБЩИЙ ШАБЛОН РУНЫ
	//
	// шаблон руны состоит из :
	//  - имени "N" какому символу соответствует (все кроме "\n")
	//  - высоты "H" равной количеству линий из котрых состоит символ
	// для примера можно посмотреть "fonts/super_fonts/style_2.txt"
	//
	// rune params | N H
	//    line 0   | o----------o
	//    line 1   | |xxxxxxxxxx|
	//    line 2   | |xxxxxxxxxx|
	//    line 3   | |xxxxxxxxxx|
	//    line 4   | |xxxxxxxxxx|
	//    line 5   | |xxxxxxxxxx|
	//         ... |      ...
	//    line n-1 | |xxxxxxxxxx|
	//    line n   | o----------o

	// для проверки того что загрузилось мы будем использовать функцию
	// lascii.GetFontNames()
	// и печатать имена загруженых шрифтов с помощью самодельной функции
	// checkFonts()

	// по мере изучения тебе придётся самому учавствовать в подготовки примера
	// для этого в коде сделаны "заготовки" с уникальными именами
	// вот пример того как они выглядят
	//
	// code for comment/uncomment [<some name>] :
	// <some code>

	// давай потренируемся и подготовим код
	// раскоменти стоку "simple uncoment"
	//
	// подсказка!
	// попробуй воспользоваться поиском ctrl + F / cmd + F

	fmt.Printf("==========\n")

	// code for comment/uncomment [simple uncoment] :
	checkFonts()
	fmt.Printf("\nBefore /\\\n")

	fmt.Printf("==========\n")

	// разберём каждый параметр по отдельности
	// чтобы увидить каким он может быть
	// и что нужно указать в файле шрифта
	// перейдите в нужную функцию для разбора
	loadSettings := lascii.LoadingSettings_t{
		Path:      path(),
		Width:     width(),
		Height:    height(),
		HardASCII: hardASCII(),
	}

	// ну и сама загрузка
	if err := lascii.LoadFonts(loadSettings); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("After  \\/\n\n")
	checkFonts()
	fmt.Printf("==========\n")

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

func path() string {
	var resultPath string

	resultPath = "defaults"
	// resultPath = "defaults/"
	// resultPath = ""

	return FONTS_DIR + "/" + resultPath
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func width() int {

	return 0
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func height() int {

	return 0
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func hardASCII() bool {

	return false
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func loadingFonts() {
	/* DESCRIPTION :
	 *
	 *   Перед использованием нужно загрузить ваши шрифты указав параметры загрузки.
	 * Шрифты можно догружать повторным вызовом LoadFonts().
	 *
	 * Path      - Путь до паки или до конкретного файла шрифта.
	 * Height    - Высота букв. Если указано меньше 1, тогда высота определяется
	 *             числом из параметров символа в файле.
	 * Width     - Ширина букв. Если указано меньше 1, тогда ширина определяется
	 *             по ширине первой строки символа из файла.
	 * HardASCII - false означает что нужно указать что это за символ.
	 *             true значит что символов в файле должно быть ровно 95 как по таблице
		 *             ascii, начиная с символа пробела.
	 *
	 * EXAMPLE COMMON CHAR CONFIG :
	 *
	 * HardASCII : false     Height : 0
	 *                  \   /
	 *                   A 6
	 *                   ████████╗
	 *                   ██╔═══██║
	 *                   ██║   ██║
	 *                   ████████║
	 *                   ██╔═══██║
	 *                   ╚═╝   ╚═╝
	*/

	loadSettings := lascii.LoadingSettings_t{
		Path:      FONTS_DIR + "/defaults",
		Width:     0,
		Height:    8,
		HardASCII: true,
	}
	if err := lascii.LoadFonts(loadSettings); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	loadSettings = lascii.LoadingSettings_t{
		Path:      FONTS_DIR + "/super_fonts/style_simple.txt",
		Width:     0,
		Height:    6,
		HardASCII: false,
	}
	if err := lascii.LoadFonts(loadSettings); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	loadSettings = lascii.LoadingSettings_t{
		Path:      FONTS_DIR + "/super_fonts/letters",
		Width:     0,
		Height:    0,
		HardASCII: false,
	}
	if err := lascii.LoadFonts(loadSettings); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	loadSettings = lascii.LoadingSettings_t{
		Path:      FONTS_DIR + "/super_fonts/shadows",
		Width:     0,
		Height:    0,
		HardASCII: false,
	}
	if err := lascii.LoadFonts(loadSettings); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
}
