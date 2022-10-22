package main

import (
	"enigma/config"
	"enigma/machine/brouilleur"
	"enigma/machine/plug_board"
	"fmt"
	"strings"
	"unicode"
)

type CommandLineInterface struct {
}

func NewCommandLineInterface() *CommandLineInterface {
	return &CommandLineInterface{}
}

func (c *CommandLineInterface) Run() error {
mainLoop:
	for {
		fmt.Printf("Please input command (info, input, change, spin, plug, unplug, cp, quit/exit): ")
		var command string
		_, _ = fmt.Scan(&command)
		switch command {
		case "info":
			fmt.Printf("Current rotor no: %v\n", brouilleur.Configuration.GetRotorNo())
			fmt.Printf("Current rotor index: %s\n", string(brouilleur.Configuration.GetRotorPositions()))
			plugs := plug_board.Configuration.Mappings()
			var pbStr []string
			for _, plug := range plugs {
				pbStr = append(pbStr, fmt.Sprintf("%c - %c", plug[0], plug[1]))
			}
			str := strings.Join(pbStr, ", ")
			if str == "" {
				str = "None"
			}
			fmt.Printf("Current plug board: %s\n", str)
			fmt.Println("OK")
		case "input":
			fmt.Printf("input mode, please input your text to encode/decode: ")
			var str string
			_, _ = fmt.Scan(&str)
			fmt.Printf("Got str: %s\n", str)
			str = strings.ToUpper(str)
			rarr := []rune(str)
			for _, r := range rarr {
				if !unicode.IsUpper(r) {
					fmt.Printf("input contains invalid char '%c', will not execute input...\n", r)
					continue mainLoop
				}
			}
			var res string
			for _, r := range rarr {
				r = plug_board.Configuration.Encode(r)
				brouilleur.Configuration.SpinRotor()
				if config.Conf.Debug {
					fmt.Printf("Current rotor index: %s\n", string(brouilleur.Configuration.GetRotorPositions()))
				}
				r = plug_board.Configuration.Encode(r)
				res += string(brouilleur.Configuration.Encode(r))
			}
			fmt.Printf("Output: %s\n", res)
			fmt.Printf("Current rotor index: %s\n", string(brouilleur.Configuration.GetRotorPositions()))
			fmt.Println("OK")
		case "change":
			fmt.Printf("change rotors, please input %d rotor numbers (seperated by space): ", config.Conf.RotorCount)
			rotorIndexes := make([]int, config.Conf.RotorCount)
			for i := 0; i < config.Conf.RotorCount; i++ {
				_, _ = fmt.Scan(&rotorIndexes[i])
			}
			fmt.Printf("Got indexes: %v\n", rotorIndexes)
			err := brouilleur.Configuration.ChangeRotors(rotorIndexes)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue mainLoop
			}
			fmt.Println("OK")
		case "spin":
			fmt.Printf("spin rotors, please input %d alphabets (altogether without space): ", config.Conf.RotorCount)
			var str string
			_, _ = fmt.Scan(&str)
			str = strings.ToUpper(str)
			fmt.Printf("Got str: %s\n", str)
			rarr := []rune(str)
			for _, r := range rarr {
				if !unicode.IsUpper(r) {
					fmt.Printf("input contains invalid char '%c', will not execute spin...\n", r)
					continue mainLoop
				}
			}
			err := brouilleur.Configuration.SetRotorPositions(rarr)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue mainLoop
			}
			fmt.Println("OK")
		case "plug":
			fmt.Printf("plug wire, please input 2 alphabets (altogether without space): ")
			var str string
			_, _ = fmt.Scan(&str)
			str = strings.ToUpper(str)
			fmt.Printf("Got str: %s\n", str)
			if len(str) != 2 || !unicode.IsUpper(rune(str[0])) || !unicode.IsUpper(rune(str[1])) {
				fmt.Printf("must provide 2 alphabets!\n")
				continue mainLoop
			}
			err := plug_board.Configuration.Plug(rune(str[0]), rune(str[1]))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue mainLoop
			}
			fmt.Println("OK")
		case "unplug":
			fmt.Printf("unplug wire, please input 2 alphabets (altogether without space): ")
			var str string
			_, _ = fmt.Scan(&str)
			str = strings.ToUpper(str)
			fmt.Printf("Got str: %s\n", str)
			if len(str) != 2 || !unicode.IsUpper(rune(str[0])) || !unicode.IsUpper(rune(str[1])) {
				fmt.Printf("must provide 2 alphabets!\n")
				continue mainLoop
			}
			err := plug_board.Configuration.Unplug(rune(str[0]), rune(str[1]))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue mainLoop
			}
			fmt.Println("OK")
		case "cp":
			fmt.Printf("Clearing plugs...")
			plug_board.Configuration.Clear()
			fmt.Printf("OK\n")
			fmt.Println()
		case "quit", "exit":
			fmt.Println("Exiting program gracefully")
			return nil
		default:
			fmt.Println("Invalid command...")
		}
	}
}
