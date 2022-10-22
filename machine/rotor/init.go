package rotor

import (
	"enigma/config"
	"fmt"
)

var (
	Configurations map[int]*Rotor // The system initializes with some predefined rotors
)

func init() {
	Configurations = make(map[int]*Rotor)
	for _, rotor := range config.Conf.Rotors {
		if rotor == nil {
			continue
		}
		if _, ok := Configurations[rotor.No]; ok {
			panic(fmt.Sprintf("got coflict NO %d", rotor.No))
		}
		r, err := NewRotorFromConf(rotor.No, rotor.Mapping)
		if err != nil {
			panic(fmt.Errorf("invalid rotor conf (no=%d): %w", rotor.No, err))
		}
		Configurations[rotor.No] = r
	}
	if len(Configurations) < config.Conf.RotorCount {
		panic(fmt.Errorf("not enough Rotors, please provide at least %d", config.Conf.RotorCount))
	}
}
