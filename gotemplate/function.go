package gotemplate

import (
	"strings"

	"github.com/Kretech/xgo/word"
)

func UseFuncAll() map[string]interface{} {
	return UseFuncSets(
		StringSet,
	)
}

func UseFuncSets(funcSets ...map[string]interface{}) map[string]interface{} {
	all := make(map[string]interface{}, 32)
	for _, funcSet := range funcSets {
		for name, fun := range funcSet {
			all[name] = fun
		}
	}
	return all
}

var StringSet = map[string]interface{}{
	"ToUpper":    strings.ToUpper,
	"ToLower":    strings.ToLower,
	"LowerFirst": word.LowerFirst,
	"UpperFirst": word.UpperFirst,
}
