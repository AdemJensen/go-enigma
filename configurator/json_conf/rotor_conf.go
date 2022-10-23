package json_conf

import (
	"enigma/machine"
	"enigma/utils"
	"fmt"
	"unicode/utf8"
)

type RotorConfig struct {
	No      int               `json:"no"`
	Mapping map[string]string `json:"mapping"`
}

func (rc *RotorConfig) MakeRotor() (*machine.Rotor, error) {
	runeMapping := make(map[rune]rune)
	for keyStr, valStr := range rc.Mapping {
		if utf8.RuneCountInString(keyStr) != 1 {
			return nil, fmt.Errorf("got invalid rotor mapping key '%s'", keyStr)
		}
		if utf8.RuneCountInString(valStr) != 1 {
			return nil, fmt.Errorf("got invalid rotor mapping val '%s'", valStr)
		}
		key := []rune(keyStr)[0]
		val := []rune(valStr)[0]
		runeMapping[key] = val
	}
	return machine.NewRotorFromMapping(rc.No, runeMapping)
}

func NewRandomRotorConfig(no int) *RotorConfig {
	return &RotorConfig{
		No:      no,
		Mapping: utils.RandomAlphabetMappingTable(),
	}
}
