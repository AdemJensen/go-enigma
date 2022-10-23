package main

import (
	"enigma/config"
	"enigma/machine"
	"enigma/machine/plug_board"
	"enigma/machine/reflector"
	"enigma/machine/rotor"
	"fmt"
)

func main() {
	var rotors []*rotor.Rotor
	for _, rIndex := range config.Conf.InitialRotors {
		r, ok := rotor.Configurations[rIndex]
		if !ok {
			panic(fmt.Errorf("rotor index %d not found", rIndex))
		}
		rotors = append(rotors, r)
	}

	mac := &machine.EnigmaMachine{
		PlugBoard: plug_board.Configuration,
		Rotors:    rotors,
		Reflector: reflector.Configuration,
	}

	app := NewCommandLineInterface(mac)
	err := app.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
