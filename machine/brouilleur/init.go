package brouilleur

import (
	"enigma/config"
	"enigma/machine/reflector"
	rotor2 "enigma/machine/rotor"
	"fmt"
)

var (
	Configuration *Brouilleur
)

func init() {
	// just select 3 rotors randomly
	var rotors []*rotor2.Rotor
	for _, rIndex := range config.Conf.InitialRotors {
		r, ok := rotor2.Configurations[rIndex]
		if !ok {
			panic(fmt.Errorf("rotor index %d not found", rIndex))
		}
		rotors = append(rotors, r)
	}
	Configuration = &Brouilleur{
		reflector: reflector.Configuration,
		rotors:    rotors,
	}
}
