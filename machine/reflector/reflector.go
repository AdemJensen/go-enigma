package reflector

import (
	"enigma/config"
	"enigma/utils"
	"fmt"
	"unicode/utf8"
)

func NewReflectorFromConf(m map[string]string) (*Reflector, error) {
	mappingTable := make(map[rune]rune)
	for keyStr, valStr := range m {
		if utf8.RuneCountInString(keyStr) != 1 {
			return nil, fmt.Errorf("got invalid reflector mapping key '%s'", keyStr)
		}
		if utf8.RuneCountInString(valStr) != 1 {
			return nil, fmt.Errorf("got invalid reflector mapping val '%s'", valStr)
		}
		key := []rune(keyStr)[0]
		val := []rune(valStr)[0]
		if _, exists := mappingTable[key]; exists {
			return nil, fmt.Errorf("conflict key '%c'", key)
		}
		mappingTable[key] = val
	}
	isInKey := make(map[rune]bool)
	isInVal := make(map[rune]bool)
	// check if every key is mapped from or to
	for key, val := range mappingTable {
		isInKey[key] = true
		isInVal[val] = true
	}
	for alp := 'A'; alp <= 'Z'; alp++ {
		if !isInKey[alp] || !isInVal[alp] {
			return nil, fmt.Errorf("missing mapping for alphabet '%c'", alp)
		}
	}
	// every key must map back to each other
	for key := 'A'; key <= 'Z'; key++ {
		val := mappingTable[key]
		if mappingTable[val] != key {
			return nil, fmt.Errorf("alphabet '%c' and '%c' does not relect back to each other", key, val)
		}
	}
	return &Reflector{
		mappingTable: mappingTable,
	}, nil
}

func NewRandomReflector() *Reflector {
	var allAlphabets []rune
	for key := 'A'; key <= 'Z'; key++ {
		allAlphabets = append(allAlphabets, key)
	}
	mappingKeys := utils.RandomSelectRuneArray(allAlphabets, 13)
	val := utils.DeleteFromRuneSlice(allAlphabets, mappingKeys...)
	mappingTable := make(map[rune]rune)
	for idx, key := range mappingKeys {
		mappingTable[key] = val[idx]
	}
	return &Reflector{
		mappingTable: mappingTable,
	}
}

func NewDefaultReflector() *Reflector {
	mapping := make(map[rune]rune)
	for key, val := 'A', 'Z'; key <= 'Z'; key, val = key+1, val-1 {
		mapping[key] = val
	}
	return &Reflector{
		mappingTable: mapping,
	}
}

// Reflector reflects 1 alphabet with another
type Reflector struct {
	mappingTable map[rune]rune
}

func (r *Reflector) Encode(v rune) rune {
	return r.mappingTable[v]
}

func (r *Reflector) ExportConf() config.ReflectorConf {
	return utils.RuneMapToStringMap(r.mappingTable)
}
