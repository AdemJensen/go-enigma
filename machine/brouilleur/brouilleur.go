package brouilleur

import (
	"enigma/config"
	reflector2 "enigma/machine/reflector"
	rotor2 "enigma/machine/rotor"
	"fmt"
	"unicode"
)

type Brouilleur struct {
	reflector *reflector2.Reflector
	rotors    []*rotor2.Rotor
}

func (b *Brouilleur) ChangeRotors(nos []int) error {
	if len(nos) != config.Conf.RotorCount {
		return fmt.Errorf("rotor must have exactly %d", config.Conf.RotorCount)
	}
	b.rotors = nil
	for _, no := range nos {
		if rot, ok := rotor2.Configurations[no]; !ok {
			return fmt.Errorf("rotor no %d not exist", no)
		} else {
			_ = rot.SetPosition('A')
			b.rotors = append(b.rotors, rot)
		}
	}
	return nil
}

func (b *Brouilleur) SetRotorPositions(pos []rune) error {
	if len(pos) != config.Conf.RotorCount {
		return fmt.Errorf("must proivde exactly %d indexes", config.Conf.RotorCount)
	}
	for _, r := range pos {
		if !unicode.IsUpper(r) {
			return fmt.Errorf("invalid index '%c'", r)
		}
	}
	for i, r := range b.rotors {
		_ = r.SetPosition(pos[i])
	}
	return nil
}

func (b *Brouilleur) GetRotorPositions() []rune {
	var res []rune
	for _, r := range b.rotors {
		res = append(res, r.CurrentPosition())
	}
	return res
}

func (b *Brouilleur) GetRotorNo() []int {
	var res []int
	for _, r := range b.rotors {
		res = append(res, r.No())
	}
	return res
}

func (b *Brouilleur) SpinRotor() {
	// spin from right to left
	for i := len(b.rotors) - 1; i >= 0; i-- {
		circulated := b.rotors[i].Spin()
		if !circulated {
			break
		}
	}
}

func (b *Brouilleur) Encode(v rune) rune {
	if config.Conf.Debug {
		fmt.Printf("%c", v)
	}
	// 1. from right to left
	for i := len(b.rotors) - 1; i >= 0; i-- {
		v = b.rotors[i].EncodeRightToLeft(v)
		if config.Conf.Debug {
			fmt.Printf("-->%c", v)
		}

	}
	// 2. reflector
	v = b.reflector.Encode(v)
	if config.Conf.Debug {
		fmt.Printf("-(reflect)->%c", v)
	}
	// 3. from left to right
	for i := 0; i < len(b.rotors); i++ {
		v = b.rotors[i].EncodeLeftToRight(v)
		if config.Conf.Debug {
			fmt.Printf("-->%c", v)
		}
	}
	if config.Conf.Debug {
		fmt.Printf("\n")
	}
	return v
}
