package machine

import (
	"fmt"
	"unicode"
)

// Rotor encodes given alphabet, and rotates itself on every keyboard press.
type Rotor struct {
	no             int
	spinOffset     int
	ringOffset     int // on which offset will this ring rotates the next rotor
	mappingTableRL map[rune]rune
	mappingTableLR map[rune]rune
}

// NewRotorFromMapping receives rotor no and mapping from right to left.
func NewRotorFromMapping(no int, ring rune, m map[rune]rune) (*Rotor, error) {
	mappingTableRL := make(map[rune]rune)
	mappingTableLR := make(map[rune]rune)
	for key, val := range m {
		if !unicode.IsUpper(key) || !unicode.IsUpper(val) {
			return nil, fmt.Errorf("invalid mapping key or val")
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
		ringOffset:     int(ring - 'A'),
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
		ringOffset:     25,
		mappingTableRL: mapping,
		mappingTableLR: mapping,
	}
}

func (r *Rotor) No() int {
	return r.no
}

// Spin returns whether to spin the next rotor
func (r *Rotor) Spin() bool {
	before := r.spinOffset
	r.spinOffset = (r.spinOffset + 1) % 26
	return before == r.ringOffset // true means should spin next
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

// SetRing change rotor's ring setting
func (r *Rotor) SetRing(v rune) error {
	if !unicode.IsUpper(v) {
		return fmt.Errorf("provided index is not an proper ring setting")
	}
	r.ringOffset = int(v - 'A')
	return nil
}

// GetRing get rotor's current ring setting
func (r *Rotor) GetRing() rune {
	return rune('A' + r.ringOffset)
}

func (r *Rotor) EncodeRightToLeft(v rune) rune {
	entryLoc := (v-'A'+rune(r.spinOffset))%26 + 'A'
	return (r.mappingTableRL[entryLoc]-'A'-rune(r.spinOffset)+26)%26 + 'A'
}

func (r *Rotor) EncodeLeftToRight(v rune) rune {
	entryLoc := (v-'A'+rune(r.spinOffset))%26 + 'A'
	return (r.mappingTableLR[entryLoc]-'A'-rune(r.spinOffset)+26)%26 + 'A'
}
