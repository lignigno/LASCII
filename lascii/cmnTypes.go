package lascii

// _________________________________________________________________________________PUBLIC

// Point in some 2D space
//
//	X --> coordinate on X axis
//	Y --> coordinate on Y axis
type Vec2_t struct {
	X int
	Y int
}

// Color using RGB scheme
//
//	R --> Red channel
//	G --> Green channel
//	B --> Blue channel
type Color_t struct {
	R int
	G int
	B int
}

type Findings_t struct {
	fontName string
	text     string
	coverage int
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

type LoadingSettings_t struct {
	Path      string
	Height    int
	Width     int
	HardASCII bool
}

// TODO : add empty runes, back rune and null rune
type FontMixSettings_t struct {
	NewName    string
	BaseName   string
	ShadowName string
	Offset     Vec2_t
}

type PrintSettings_t struct {
	Align       uint8
	LineSpace   uint8
	LetterSpace uint8
	MaxWidth    int
}

// sgr param from SGR_BOLD to SGR_BACK_WHITE_HI
type SGRParam_t = uint8

// more then 0 or not, maybe converted from color
type SGRValue_t = uint32

// SGR_COLOR_MODE_STANDART or SGR_COLOR_MODE_EXTENDED
type SGRColorMode_t = uint8

// settings for each SGR group
type SGRSettings_t [_NUM_SGR_GROUPS]SGRValue_t

type LSB_t struct {
	L SGRSettings_t // L - LETTER
	S SGRSettings_t // S - SHADOW
	B SGRSettings_t // B - BACK
}

// ________________________________________________________________________________PRIVATE

// Sizes for 2D objects
//
//	W --> Width of objects
//	H --> Height of objects
type _Size2_t struct {
	W int
	H int
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

// Minimum data for the description of the symbol
//
//	Size --> Size of the symbol image
//	Symbol --> 2D array for runes of symbol image
type _BaseRune_t struct {
	Size  _Size2_t
	Runes [][]rune
}

type _Rune_t struct {
	Letter _BaseRune_t
	Shadow *_BaseRune_t
}

type _Font_t struct {
	EmptyRuneL rune
	EmptyRuneS rune
	BackRune   rune
	Offset     Vec2_t
	NullRune   *_Rune_t
	MaxSize    _Size2_t
	Runes      map[rune]_Rune_t
}

type _FontsLib_t struct {
	// maybe other params some "rune and what font is the rune from?"
	Fonts map[string]_Font_t
}

type _SGRGroup_t struct {
	Params []SGRParam_t
	Unset  SGRParam_t
}

type _FileDestination_t struct {
	Dir      string
	FileName string
	Ext      string
}

type _LSBTabe_t = [_NUM_CHANGES_VARIANTS]string

type _BluePrint_t struct {
	Size         _Size2_t
	LettersPos   []Vec2_t
	RowCanvas    [][]_CanvasCell
	RealWidths   []int
	Changes      _LSBTabe_t
	ChangesExist bool
}

type _DetLayer = int32

// Cell of row canvas
//
//	Det --> determinant of rune layer (L, S or B)
//	R --> some rune of image
type _CanvasCell struct {
	Det _DetLayer
	R   rune
}
