package machine

import (
	"fmt"
)

func NewReflectorFromMapping(mappingTable map[rune]rune) (*Reflector, error) {
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
