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
	// в этом примере мы не будем печатать сам текст
	// толко изучать как можно загружать шрифты

	// ОБЩИЙ ШАБЛОН РУНЫ
	//
	// шаблон руны состоит из :
	//  - имени "N" какому символу соответствует (все кроме "\n")
	//  - высоты "H" равной количеству линий из котрых состоит символ
	// для примера можно посмотреть "fonts/super_fonts/style_2.txt"
	// все параметры в rune params опциональны и могут отсутствовать
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
	// раскомментируй стоку "simple uncoment"
	//
	// подсказка!
	// попробуй воспользоваться поиском ctrl + F / cmd + F

	fmt.Printf("==========\n")

	// code for comment/uncomment [simple uncoment] :
	// checkFonts()
	fmt.Printf("\nBefore /\\\n")

	fmt.Printf("==========\n")

	// разберём каждый параметр по отдельности
	// чтобы увидить каким он может быть
	// и что нужно указать в файле шрифта
	// перейди в нужную функцию для разбора
	loadSettings := lascii.LoadingSettings_t{
		Path:      path(),
		Width:     width(),
		Height:    height(),
		HardASCII: hardASCII(),
	}

	fmt.Printf("loadSettings\n\n")
	fmt.Printf(" - Path      : %s\n", loadSettings.Path)
	fmt.Printf(" - Width     : %d\n", loadSettings.Width)
	fmt.Printf(" - Height    : %d\n", loadSettings.Height)
	fmt.Printf(" - HardASCII : %v\n", loadSettings.HardASCII)

	fmt.Printf("==========\n")

	// ну и сама загрузка
	if err := lascii.LoadFonts(loadSettings); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	// ещё немного для любопытных
	// loadingFonts()

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
	resultPath := ""

	// чтобы загрузить шрифты можно указать папку или фаил
	// у нас путь будет состоять из двух переменных
	//
	// FONTS_DIR - путь до папки со всеми шрифтами
	// resultPath - путь внутри папки FONTS_DIR

	// давайте для начала попробуем с указанием каких-то дефолтных шрифтов
	//
	// раскоментируй строку "path ex0",
	// запусти и посмотри что получится

	// code for comment/uncomment [path ex0] :
	// resultPath = "defaults/"

	// добавились шрифты которые лежали в папке "defaults"
	// к слову папка может быть указана без символа "/"

	// code for comment/uncomment [path ex1] :
	// resultPath = "defaults"

	// если в папке нету файлов то ничего не будет загружено

	// code for comment/uncomment [path ex2] :
	// resultPath = ""

	// если в папе будут находится файлы не с шрифтами то будет выдана ошибка

	// code for comment/uncomment [path ex3] :
	// resultPath = "error/0"

	// чтобы загружались не все файлы в папке то можно указвать их по отдельности

	// code for comment/uncomment [path ex4] :
	// resultPath = "error/0/valid"

	// остальные настраиваемые параметры загрузки (о каждом из них далее)
	// распростроняются на весь указанный путь
	// это значит что если в директории лежат несколько разных файлов с шрифтами
	// и если настройки загрузки подходят для одних, но не подходят для других
	// то будет ошибка

	// code for comment/uncomment [path ex5] :
	// resultPath = "error/1"

	// code for comment/uncomment [path ex6] :
	// resultPath = "otherHeight"

	// code for comment/uncomment [path ex7] :
	// resultPath = "super_fonts/style_simple.txt"

	return FONTS_DIR + "/" + resultPath
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func width() int {
	resultWidth := 0

	// есть два варианта ширины
	//         /              \
	//        /                \
	// автоматическая        строгая
	//    val < 1            val >= 1

	// val < 1 означает что ширина определяется автоматически
	// по ширине первой линии (line 0) и все последующие строки руны
	// должны совпадать по ширине с первой

	// если указать val >= 1 то ширна всех рун должна равнятся val

	// давай сделаем неверные настройки для шрифтов в папке "defaults"
	// закомментируй :
	// "path ex1"
	// "path ex2"
	// "path ex3"
	// "path ex4"
	// "path ex5"
	// раскоментируй :
	// "width ex0"

	// code for comment/uncomment [width ex0] :
	// resultWidth = 6

	// 6 это ширина первого символа " "
	// так как обычно используется разная ширина для символов, то закоментируй "width ex0"

	return resultWidth
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func height() int {
	resultHeight := 8

	// здесь всё также как и с шириной
	// автоматическая и строгая
	// но при автоматической высота берётся из строки "rune params" общего шаблона

	// давай попробуем загрузить шрифт super_space
	// раскоментируй "path ex6" и "height ex0"
	// и найди сам шрифт чтобы посмотреть что в нём указано

	// code for comment/uncomment [height ex0] :
	// resultHeight = 0

	// code for comment/uncomment [height ex1] :
	// resultHeight = 6

	return resultHeight
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func hardASCII() bool {
	resultHardASCII := true

	// данный параметр отвечает за то из каких символов будет состоять загруженный шрифт

	// если true то это должны быть символы из аски таблицы строго по порядку как в таблице
	// все c " " до "~" включительно

	// если "hardASCII" будет false то в строке "rune params"
	// нужно первым указывать руну к которой относится картинка

	// для примера можно попробовать загрузить шрифт style_simple.txt
	// раскоментируй
	// "path ex7"
	// "height ex1"
	// "hardASCII ex0"

	// code for comment/uncomment [hardASCII ex0] :
	// resultHardASCII = false

	return resultHardASCII
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func loadingFonts() {
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

	// если ты всётаки добрався до сюда то смотри
	// повторная загрузка ничего не добавит
	loadSettings = lascii.LoadingSettings_t{
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
