package main

import (
	"enigma/cli"
	"enigma/configurator/json_conf"
	"enigma/machine"
	"fmt"
)

func main() {
	cnf, err := json_conf.NewConfigFromFile("./conf.json")
	if err != nil {
		panic(fmt.Errorf("read config file error: %w", err))
	}

	rotorObjs, err := cnf.MakeInitialRotors()
	if err != nil {
		panic(fmt.Errorf("make rotors failed: %w", err))
	}

	reflectorObj, err := cnf.Reflector.MakeReflector()
	if err != nil {
		panic(fmt.Errorf("make reflector failed: %w", err))
	}

	plgBoardObj, err := cnf.MakeInitialPlugBoard()
	if err != nil {
		panic(fmt.Errorf("make plug board failed: %w", err))
	}

	mac := &machine.EnigmaMachine{
		PlugBoard: plgBoardObj,
		Rotors:    rotorObjs,
		Reflector: reflectorObj,
	}

	app := cli.NewCommandLineInterface(mac, cnf)
	if err := app.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
