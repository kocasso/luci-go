// Code generated by "stringer -type=ComponentType"; DO NOT EDIT.

package ui

import "strconv"

const _ComponentType_name = "RecipeStepSummary"

var _ComponentType_index = [...]uint8{0, 6, 10, 17}

func (i ComponentType) String() string {
	if i < 0 || i >= ComponentType(len(_ComponentType_index)-1) {
		return "ComponentType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ComponentType_name[_ComponentType_index[i]:_ComponentType_index[i+1]]
}
