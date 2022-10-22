package plug_board

import (
	"enigma/config"
	"fmt"
	"unicode/utf8"
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

func (p *PlugBoard) ApplyConfig(cfg []config.PlugConfig) error {
	for i, c := range cfg {
		if len(c) != 2 {
			return fmt.Errorf("plug config must contain 2 strings for every set")
		}
		if utf8.RuneCountInString(c[0]) != 1 {
			return fmt.Errorf("got invalid plug board set[0] key '%s'", c[0])
		}
		if utf8.RuneCountInString(c[1]) != 1 {
			return fmt.Errorf("got invalid plug board set[1] val '%s'", c[1])
		}
		c1 := []rune(c[0])[0]
		c2 := []rune(c[1])[0]
		err := p.Plug(c1, c2)
		if err != nil {
			return fmt.Errorf("cannot process plug conf %d: %v", i, err)
		}
	}
	return nil
}

func (p *PlugBoard) Encode(v rune) rune {
	return p.switchMapping[v]
}

func (p *PlugBoard) Plug(a, b rune) error {
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
	if !p.occupiedPlugs[a] || !p.occupiedPlugs[b] || p.switchMapping[a] != b || p.switchMapping[b] != a {
		return fmt.Errorf("plug (%c - %c) are not linked together", a, b)
	}
	p.occupiedPlugs[a] = false
	p.occupiedPlugs[b] = false
	p.switchMapping[a] = a
	p.switchMapping[b] = b
	return nil
}

func (p *PlugBoard) Clear() {
	mappings := p.Mappings()
	for _, m := range mappings {
		_ = p.Unplug(m[0], m[1])
	}
}

func (p *PlugBoard) Mappings() [][]rune {
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
