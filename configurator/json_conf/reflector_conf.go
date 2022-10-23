package json_conf

import (
	"enigma/machine"
	"enigma/utils"
	"fmt"
	"unicode/utf8"
)

type ReflectorConf map[string]string

func (rc ReflectorConf) MakeReflector() (*machine.Reflector, error) {
	if rc == nil {
		return nil, fmt.Errorf("got nil config")
	}
	runeMapping := make(map[rune]rune)
	for keyStr, valStr := range rc {
		if utf8.RuneCountInString(keyStr) != 1 {
			return nil, fmt.Errorf("got invalid reflector mapping key '%s'", keyStr)
		}
		if utf8.RuneCountInString(valStr) != 1 {
			return nil, fmt.Errorf("got invalid reflector mapping val '%s'", valStr)
		}
		key := []rune(keyStr)[0]
		val := []rune(valStr)[0]
		runeMapping[key] = val
	}
	return machine.NewReflectorFromMapping(runeMapping)
}

func NewRandomReflectorConfig() ReflectorConf {
	var allAlphabets []rune
	for key := 'A'; key <= 'Z'; key++ {
		allAlphabets = append(allAlphabets, key)
	}
	mappingKeys := utils.RandomSelectRuneArray(allAlphabets, 13)
	val := utils.DeleteFromRuneSlice(allAlphabets, mappingKeys...)
	mappingTable := make(map[string]string)
	for idx, key := range mappingKeys {
		mappingTable[string(key)] = string(val[idx])
		mappingTable[string(val[idx])] = string(key)
	}
	return mappingTable
}
