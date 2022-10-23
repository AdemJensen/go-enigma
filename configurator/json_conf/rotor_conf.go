package json_conf

import (
	"enigma/machine"
	"enigma/utils"
	"fmt"
	"unicode/utf8"
)

type RotorConfig struct {
	No          int               `json:"no"`
	RingSetting string            `json:"ring_setting"`
	Mapping     map[string]string `json:"mapping"`
}

func (rc *RotorConfig) MakeRotor() (*machine.Rotor, error) {
	if utf8.RuneCountInString(rc.RingSetting) != 1 {
		return nil, fmt.Errorf("got invalid ring setting '%s'", rc.RingSetting)
	}
	ringSetting := []rune(rc.RingSetting)[0]
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
	return machine.NewRotorFromMapping(rc.No, ringSetting, runeMapping)
}

func NewRandomRotorConfig(no int) *RotorConfig {
	return &RotorConfig{
		No:          no,
		RingSetting: string(utils.RandomAlphabet()),
		Mapping:     utils.RandomAlphabetMappingTable(),
	}
}
