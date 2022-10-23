package json_conf

import (
	"enigma/machine"
	"fmt"
	"unicode/utf8"
)

type PlugConfig []string

func (pc PlugConfig) ApplyTo(plg *machine.PlugBoard) error {
	if len(pc) != 2 {
		return fmt.Errorf("plug config must contain 2 strings for every set")
	}
	if utf8.RuneCountInString(pc[0]) != 1 {
		return fmt.Errorf("got invalid plug board set[0] key '%s'", pc[0])
	}
	if utf8.RuneCountInString(pc[1]) != 1 {
		return fmt.Errorf("got invalid plug board set[1] val '%s'", pc[1])
	}
	c1 := []rune(pc[0])[0]
	c2 := []rune(pc[1])[0]
	err := plg.Plug(c1, c2)
	if err != nil {
		return fmt.Errorf("apply plug conf error: %v", err)
	}
	return nil
}
