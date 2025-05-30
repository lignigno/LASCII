package lascii

// ____________________________________________________________________________RESIZE RUNE

func resizeRune(original _BaseRune_t, newSize _Size2_t) _BaseRune_t {
	newRune := _BaseRune_t{Size: newSize}

	newRune.Runes = make([][]rune, newSize.H)
	for y := 0; y < newRune.Size.H; y++ {
		newRune.Runes[y] = make([]rune, newSize.W)
	}

	stepX := float32(original.Size.W) / float32(newSize.W)
	stepY := float32(original.Size.H) / float32(newSize.H)

	for y := range newRune.Size.H {
		for x := range newRune.Size.W {
			newRune.Runes[y][x] = original.Runes[int(float32(y)*stepY)][int(float32(x)*stepX)]
		}
	}

	return newRune
}
