package rotor

import (
	"enigma/config"
	"enigma/utils"
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Rotor encodes given alphabet, and rotates itself on every keyboard press.
type Rotor struct {
	no             int
	spinOffset     int
	mappingTableRL map[rune]rune
	mappingTableLR map[rune]rune
}

func NewRotorFromConf(no int, m map[string]string) (*Rotor, error) {
	mappingTableRL := make(map[rune]rune)
	mappingTableLR := make(map[rune]rune)
	for keyStr, valStr := range m {
		if utf8.RuneCountInString(keyStr) != 1 {
			return nil, fmt.Errorf("got invalid rotor mapping key '%s'", keyStr)
		}
		if utf8.RuneCountInString(valStr) != 1 {
			return nil, fmt.Errorf("got invalid rotor mapping val '%s'", valStr)
		}
		key := []rune(keyStr)[0]
		val := []rune(valStr)[0]
		if _, exists := mappingTableRL[key]; exists {
			return nil, fmt.Errorf("conflict key '%c'", key)
		}
		mappingTableRL[key] = val
		mappingTableLR[val] = key
	}
	isInKey := make(map[rune]bool)
	isInVal := make(map[rune]bool)
	// check if every key is mapped from or to
	for key, val := range mappingTableRL {
		isInKey[key] = true
		isInVal[val] = true
	}
	for alp := 'A'; alp <= 'Z'; alp++ {
		if !isInKey[alp] || !isInVal[alp] {
			return nil, fmt.Errorf("missing mapping for alphabet '%c'", alp)
		}
	}
	return &Rotor{
		no:             no,
		spinOffset:     0,
		mappingTableRL: mappingTableRL,
		mappingTableLR: mappingTableLR,
	}, nil
}

func NewDefaultRotor(no int) *Rotor {
	mapping := make(map[rune]rune)
	for alp := 'A'; alp <= 'Z'; alp++ {
		mapping[alp] = alp
	}
	return &Rotor{
		no:             no,
		spinOffset:     0,
		mappingTableRL: mapping,
		mappingTableLR: mapping,
	}
}

func NewRandomRotor(no int) *Rotor {
	rot, err := NewRotorFromConf(no, utils.RuneMapToStringMap(utils.RandomAlphabetMappingTable()))
	if err != nil {
		panic(err)
	}
	return rot
}

func (r *Rotor) No() int {
	return r.no
}

func (r *Rotor) Spin() bool {
	before := r.spinOffset
	r.spinOffset = (r.spinOffset + 1) % 26
	return before > r.spinOffset // true means circulated
}

func (r *Rotor) SetPosition(v rune) error {
	if !unicode.IsUpper(v) {
		return fmt.Errorf("provided index is not an upper alphabet")
	}
	r.spinOffset = int(v - 'A')
	return nil
}

func (r *Rotor) CurrentPosition() rune {
	return rune('A' + r.spinOffset)
}

func (r *Rotor) EncodeRightToLeft(v rune) rune {
	entryLoc := (v-'A'+rune(r.spinOffset))%26 + 'A'
	return (r.mappingTableRL[entryLoc]-'A'-rune(r.spinOffset)+26)%26 + 'A'
}

func (r *Rotor) EncodeLeftToRight(v rune) rune {
	entryLoc := (v-'A'+rune(r.spinOffset))%26 + 'A'
	return (r.mappingTableLR[entryLoc]-'A'-rune(r.spinOffset)+26)%26 + 'A'
}

func (r *Rotor) ExportConf() *config.RotorConfig {
	return &config.RotorConfig{
		No:      r.no,
		Mapping: utils.RuneMapToStringMap(r.mappingTableRL),
	}
}
