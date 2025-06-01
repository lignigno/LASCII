package lascii

// ______________________________________________________________________________VARIABLES

var _fontsLib _FontsLib_t = _FontsLib_t{
	Fonts: make(map[string]_Font_t),
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

var _printSettings PrintSettings_t = PrintSettings_t{
	Align:       ALIGN_LEFT,
	LineSpace:   0,
	LetterSpace: 0,
	MaxWidth:    200,
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

var _sgrGroups [_NUM_SGR_GROUPS]_SGRGroup_t = [_NUM_SGR_GROUPS]_SGRGroup_t{
	{
		Params: []SGRParam_t{
			_SGR_RESET_ALL,
			_SGR_RESET_ALL,
		},
		Unset: _SGR_UNSET_ALL,
	},
	{
		Params: []SGRParam_t{
			SGR_FONT_BLACK,
			SGR_FONT_COLOR,
			SGR_FONT_BLACK_HI,
			SGR_FONT_WHITE_HI,
		},
		Unset: _SGR_UNSET_GROUP_FONT_COLOR,
	},
	{
		Params: []SGRParam_t{
			SGR_BACK_BLACK,
			SGR_BACK_COLOR,
			SGR_BACK_BLACK_HI,
			SGR_BACK_WHITE_HI,
		},
		Unset: _SGR_UNSET_GROUP_BACK_COLOR,
	},
	{
		Params: []SGRParam_t{
			SGR_BOLD,
			SGR_FADED,
		},
		Unset: _SGR_UNSET_GROUP_BOLD_FADED,
	},
	{
		Params: []SGRParam_t{
			SGR_ITALICS,
			SGR_ITALICS,
		},
		Unset: _SGR_UNSET_GROUP_ITALICS,
	},
	{
		Params: []SGRParam_t{
			SGR_UNDERLINE,
			SGR_UNDERLINE,
		},
		Unset: _SGR_UNSET_GROUP_UNDERLINE,
	},
	{
		Params: []SGRParam_t{
			SGR_BLINK_SLOW,
			SGR_BLINK_FAST,
		},
		Unset: _SGR_UNSET_GROUP_BLINK,
	},
	{
		Params: []SGRParam_t{
			SGR_NEGATIVE,
			SGR_NEGATIVE,
		},
		Unset: _SGR_UNSET_GROUP_NEGATIVE,
	},
	{
		Params: []SGRParam_t{
			SGR_HIDDEN,
			SGR_HIDDEN,
		},
		Unset: _SGR_UNSET_GROUP_HIDDEN,
	},
	{
		Params: []SGRParam_t{
			SGR_STRIKETHROUGH,
			SGR_STRIKETHROUGH,
		},
		Unset: _SGR_UNSET_GROUP_STRIKETHROUGH,
	},
	{
		Params: []SGRParam_t{
			SGR_OVERLINE,
			SGR_OVERLINE,
		},
		Unset: _SGR_UNSET_GROUP_OVERLINE,
	},
	{
		Params: []SGRParam_t{
			SGR_ALTERNATIVE_1,
			SGR_FRACTURE,
		},
		Unset: _SGR_UNSET_GROUP_STILE,
	},
	{
		Params: []SGRParam_t{
			SGR_FRAMED,
			SGR_SURROUNDED,
		},
		Unset: _SGR_UNSET_GROUP_DECORATION,
	},
	{
		Params: []SGRParam_t{
			SGR_IDEOGRAM_UNDERLINE,
			SGR_IDEOGRAM_STRESS_MARK,
		},
		Unset: _SGR_UNSET_GROUP_IDEOGRAM,
	},
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

var _baseRuneNull _BaseRune_t = _BaseRune_t{
	Size: _Size2_t{W: 5, H: 5},
	Runes: [][]rune{
		[]rune("o---o"),
		[]rune("|. .|"),
		[]rune("| . |"),
		[]rune("|. .|"),
		[]rune("o---o"),
	},
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

var _runeNull _Rune_t = _Rune_t{
	Letter: _baseRuneNull,
	Shadow: &_baseRuneNull,
}
