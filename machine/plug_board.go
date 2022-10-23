package machine

import (
	"enigma/errors"
	"enigma/validations"
	"fmt"
)

type PlugBoard struct {
	occupiedPlugs map[rune]bool
	switchMapping map[rune]rune
}

func NewPlugBoard() *PlugBoard {
	mappingTable := make(map[rune]rune)
	for key := 'A'; key <= 'Z'; key++ {
		mappingTable[key] = key
	}
	return &PlugBoard{
		occupiedPlugs: make(map[rune]bool),
		switchMapping: mappingTable,
	}
}

func (p *PlugBoard) Encode(v rune) rune {
	return p.switchMapping[v]
}

func (p *PlugBoard) Plug(a, b rune) error {
	if !validations.IsValidPlugBoardPin(a) {
		return errors.NewInputError(errors.InputTypePlugBoardPin, string(a))
	}
	if !validations.IsValidPlugBoardPin(b) {
		return errors.NewInputError(errors.InputTypePlugBoardPin, string(b))
	}
	if a == b {
		return fmt.Errorf("cannot plug onto same pin (%c)", a)
	}
	if p.occupiedPlugs[a] {
		return fmt.Errorf("plug %c has been occupied (%c - %c)", a, a, p.switchMapping[a])
	}
	if p.occupiedPlugs[b] {
		return fmt.Errorf("plug %c has been occupied (%c - %c)", b, b, p.switchMapping[b])
	}
	p.occupiedPlugs[a] = true
	p.occupiedPlugs[b] = true
	p.switchMapping[a] = b
	p.switchMapping[b] = a
	return nil
}

func (p *PlugBoard) Unplug(a, b rune) error {
	if !validations.IsValidPlugBoardPin(a) {
		return errors.NewInputError(errors.InputTypePlugBoardPin, string(a))
	}
	if !validations.IsValidPlugBoardPin(b) {
		return errors.NewInputError(errors.InputTypePlugBoardPin, string(b))
	}
	if a == b {
		return fmt.Errorf("cannot unplug same pin (%c)", a)
	}
	if !p.occupiedPlugs[a] || !p.occupiedPlugs[b] || p.switchMapping[a] != b || p.switchMapping[b] != a {
		return fmt.Errorf("plug (%c - %c) are not linked together", a, b)
	}
	p.occupiedPlugs[a] = false
	p.occupiedPlugs[b] = false
	p.switchMapping[a] = a
	p.switchMapping[b] = b
	return nil
}

// IsLinked returns whether the 2 pins are linked or not.
// If input same plug, you can use to tell whether the pin is occupied.
func (p *PlugBoard) IsLinked(a, b rune) (bool, error) {
	if !validations.IsValidPlugBoardPin(a) {
		return false, errors.NewInputError(errors.InputTypePlugBoardPin, string(a))
	}
	if !validations.IsValidPlugBoardPin(b) {
		return false, errors.NewInputError(errors.InputTypePlugBoardPin, string(b))
	}
	return p.switchMapping[a] == b, nil
}

func (p *PlugBoard) Clear() {
	mappings := p.PluggedCables()
	for _, m := range mappings {
		_ = p.Unplug(m[0], m[1])
	}
}

func (p *PlugBoard) PluggedCables() [][]rune {
	alreadyOutput := make(map[rune]bool)
	var res [][]rune
	for key, val := range p.switchMapping {
		if key == val || alreadyOutput[key] || alreadyOutput[val] {
			continue
		}
		alreadyOutput[key] = true
		alreadyOutput[val] = true
		res = append(res, []rune{key, val})
	}
	return res
}
