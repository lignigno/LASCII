package lascii

// _____________________________________________________________CONVERT COLOR TO SGR VALUE

func ConvertColor2SGRValue(mode SGRColorMode_t, color Color_t) SGRValue_t {
	return SGRValue_t(color.B<<24 | color.G<<16 | color.R<<8 | int(mode))
}

// __________________________________________________________________________SET SGR PARAM

func convetParam2Value(param SGRParam_t, value SGRValue_t) SGRValue_t {
	if param == SGR_FONT_COLOR || param == SGR_BACK_COLOR {
		return value
	}

	if value > 0 {
		return SGRValue_t(param)
	}

	return 0
}

//                                                                                       |
// --------------------------------------------------------------------------------------|
//                                                                                       |

func SetSGRParam(layer *SGRSettings_t, param SGRParam_t, value SGRValue_t) {
	for groupID := 0; groupID < _NUM_SGR_GROUPS; groupID++ {
		for i := 0; i < len(_sgrGroups[groupID].Params); i += 2 {
			from := _sgrGroups[groupID].Params[i]
			to := _sgrGroups[groupID].Params[i+1]

			if from <= param && param <= to {
				(*layer)[groupID] = convetParam2Value(param, value)
				return
			}
		}
	}
}
