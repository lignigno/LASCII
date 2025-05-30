package lascii

// _________________________________________________________________________________PUBLIC

const (
	ALIGN_LEFT = iota
	ALIGN_RIGHT
	ALIGN_CENTER
	ALIGN_JUSTIFY
)

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

const (
	SGR_BOLD                 = SGRParam_t(1)
	SGR_FADED                = SGRParam_t(2)
	SGR_ITALICS              = SGRParam_t(3)
	SGR_UNDERLINE            = SGRParam_t(4)
	SGR_BLINK_SLOW           = SGRParam_t(5)
	SGR_BLINK_FAST           = SGRParam_t(6)
	SGR_NEGATIVE             = SGRParam_t(7)
	SGR_HIDDEN               = SGRParam_t(8)
	SGR_STRIKETHROUGH        = SGRParam_t(9)
	SGR_ALTERNATIVE_1        = SGRParam_t(11)
	SGR_ALTERNATIVE_2        = SGRParam_t(12)
	SGR_ALTERNATIVE_3        = SGRParam_t(13)
	SGR_ALTERNATIVE_4        = SGRParam_t(14)
	SGR_ALTERNATIVE_5        = SGRParam_t(15)
	SGR_ALTERNATIVE_6        = SGRParam_t(16)
	SGR_ALTERNATIVE_7        = SGRParam_t(17)
	SGR_ALTERNATIVE_8        = SGRParam_t(18)
	SGR_ALTERNATIVE_9        = SGRParam_t(19)
	SGR_FRACTURE             = SGRParam_t(20)
	SGR_FONT_BLACK           = SGRParam_t(30)
	SGR_FONT_RED             = SGRParam_t(31)
	SGR_FONT_GREEN           = SGRParam_t(32)
	SGR_FONT_YELLOW          = SGRParam_t(33)
	SGR_FONT_BLUE            = SGRParam_t(34)
	SGR_FONT_MAGENTA         = SGRParam_t(35)
	SGR_FONT_CYAN            = SGRParam_t(36)
	SGR_FONT_WHITE           = SGRParam_t(37)
	SGR_FONT_COLOR           = SGRParam_t(38)
	SGR_BACK_BLACK           = SGRParam_t(40)
	SGR_BACK_RED             = SGRParam_t(41)
	SGR_BACK_GREEN           = SGRParam_t(42)
	SGR_BACK_YELLOW          = SGRParam_t(43)
	SGR_BACK_BLUE            = SGRParam_t(44)
	SGR_BACK_MAGENTA         = SGRParam_t(45)
	SGR_BACK_CYAN            = SGRParam_t(46)
	SGR_BACK_WHITE           = SGRParam_t(47)
	SGR_BACK_COLOR           = SGRParam_t(48)
	SGR_FRAMED               = SGRParam_t(51)
	SGR_SURROUNDED           = SGRParam_t(52)
	SGR_OVERLINE             = SGRParam_t(53)
	SGR_IDEOGRAM_UNDERLINE   = SGRParam_t(60)
	SGR_IDEOGRAM_D_UNDERLINE = SGRParam_t(61)
	SGR_IDEOGRAM_OVERLINE    = SGRParam_t(62)
	SGR_IDEOGRAM_D_OVERLINE  = SGRParam_t(63)
	SGR_IDEOGRAM_STRESS_MARK = SGRParam_t(64)
	SGR_FONT_BLACK_HI        = SGRParam_t(90)
	SGR_FONT_RED_HI          = SGRParam_t(91)
	SGR_FONT_GREEN_HI        = SGRParam_t(92)
	SGR_FONT_YELLOW_HI       = SGRParam_t(93)
	SGR_FONT_BLUE_HI         = SGRParam_t(94)
	SGR_FONT_MAGENTA_HI      = SGRParam_t(95)
	SGR_FONT_CYAN_HI         = SGRParam_t(96)
	SGR_FONT_WHITE_HI        = SGRParam_t(97)
	SGR_BACK_BLACK_HI        = SGRParam_t(100)
	SGR_BACK_RED_HI          = SGRParam_t(101)
	SGR_BACK_GREEN_HI        = SGRParam_t(102)
	SGR_BACK_YELLOW_HI       = SGRParam_t(103)
	SGR_BACK_BLUE_HI         = SGRParam_t(104)
	SGR_BACK_MAGENTA_HI      = SGRParam_t(105)
	SGR_BACK_CYAN_HI         = SGRParam_t(106)
	SGR_BACK_WHITE_HI        = SGRParam_t(107)
)

const (
	SGR_PARAM_OFF = SGRValue_t(iota)
	SGR_PARAM_ON
)

const (
	SGR_COLOR_MODE_STANDART = SGRColorMode_t(5)
	SGR_COLOR_MODE_EXTENDED = SGRColorMode_t(2)
)

// ________________________________________________________________________________PRIVATE

const (
	_DEFAULT_PERMISSION = 0644
	_FIRST_ASCII_CODE   = 0x20
	_LAST_ASCII_CODE    = 0x7e

	_NUM_SGR_GROUPS = 14

	_SGR_RESET_SEQUENCE = "\033[m"
)

const (
	_DET_NULL    = _DetLayer(0)
	_DET_L_RUNE  = _DetLayer(10)
	_DET_L_EMPTY = _DetLayer(11)
	_DET_S_RUNE  = _DetLayer(20)
	_DET_S_EMPTY = _DetLayer(21)
	_DET_BACK    = _DetLayer(30)
)

const (
	_SGR_RESET_ALL                 = SGRParam_t(255)
	_SGR_UNSET_ALL                 = 0
	_SGR_UNSET_GROUP_BOLD_FADED    = 22
	_SGR_UNSET_GROUP_ITALICS       = 23
	_SGR_UNSET_GROUP_UNDERLINE     = 24
	_SGR_UNSET_GROUP_BLINK         = 25
	_SGR_UNSET_GROUP_NEGATIVE      = 27
	_SGR_UNSET_GROUP_HIDDEN        = 28
	_SGR_UNSET_GROUP_STRIKETHROUGH = 29
	_SGR_UNSET_GROUP_STILE         = 10
	_SGR_UNSET_GROUP_FONT_COLOR    = 39
	_SGR_UNSET_GROUP_BACK_COLOR    = 49
	_SGR_UNSET_GROUP_DECORATION    = 54
	_SGR_UNSET_GROUP_OVERLINE      = 55
	_SGR_UNSET_GROUP_IDEOGRAM      = 65
)
