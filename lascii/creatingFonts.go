package lascii

import (
	"errors"
)

// ___________________________________________________________________________SUBFUNCTIONS

func validateSettings(settings FontMixSettings_t) error {
	if _, ok := _fontsLib.Fonts[settings.NewName]; ok {
		return errors.New("font {" + settings.NewName + "} exist")
	}
	if _, ok := _fontsLib.Fonts[settings.BaseName]; !ok {
		return errors.New("incorrect font name {" + settings.BaseName + "}")
	}
	if _, ok := _fontsLib.Fonts[settings.ShadowName]; settings.WithShadow && !ok {
		return errors.New("incorrect font name {" + settings.ShadowName + "}")
	}

	return nil
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func mixRunes(newFont *_Font_t, settings FontMixSettings_t) {
	newFont.Runes = make(map[rune]_Rune_t)

	for k, v := range _fontsLib.Fonts[settings.BaseName].Runes {
		if settings.WithShadow {
			if newFont.NullRune != nil {
				// буква есть, а тени нету тогда краказября у тени
				v.Shadow = newFont.NullRune.Shadow
			}

			if char, ok := _fontsLib.Fonts[settings.ShadowName].Runes[k]; ok {
				v.Shadow = &char.Letter
				newFont.EmptyRuneS = _fontsLib.Fonts[settings.ShadowName].EmptyRuneL
			}
		}

		newFont.Runes[k] = v
	}
}

// ______________________________________________________________________________MAIN FUNC

func CreateFontMix(settings FontMixSettings_t) error {
	newFont := _Font_t{}

	if err := validateSettings(settings); err != nil {
		return err
	}

	newFont.EmptyRuneL = _fontsLib.Fonts[settings.BaseName].EmptyRuneL
	newFont.EmptyRuneS = _fontsLib.Fonts[settings.BaseName].EmptyRuneS
	newFont.BackRune = _fontsLib.Fonts[settings.BaseName].BackRune
	newFont.NullRune = _fontsLib.Fonts[settings.BaseName].NullRune
	newFont.MaxSize = _fontsLib.Fonts[settings.BaseName].MaxSize
	newFont.Offset = settings.Offset

	if newFont.NullRune != nil {
		newNullRune := _Rune_t{}

		newNullRune.Letter = resizeRune(newFont.NullRune.Letter, newFont.MaxSize)
		tmp := resizeRune(*newFont.NullRune.Shadow, newFont.MaxSize)
		newNullRune.Shadow = &tmp

		newFont.NullRune = &newNullRune
	}

	mixRunes(&newFont, settings)
	_fontsLib.Fonts[settings.NewName] = newFont

	return nil
}
