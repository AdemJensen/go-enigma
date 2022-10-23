package cli

import (
	"enigma/configurator/json_conf"
	"enigma/machine"
	"enigma/utils"
	"fmt"
	"strings"
	"unicode"
)

type CommandLineInterface struct {
	enigmaMachine *machine.EnigmaMachine
	conf          *json_conf.Config
}

func NewCommandLineInterface(em *machine.EnigmaMachine, conf *json_conf.Config) *CommandLineInterface {
	return &CommandLineInterface{
		enigmaMachine: em,
		conf:          conf,
	}
}

func (c *CommandLineInterface) Run() error {
	if c.enigmaMachine == nil || !c.enigmaMachine.IsValid() {
		return fmt.Errorf("invalid enigma machine")
	}
mainLoop:
	for {
		fmt.Printf("Please input command (info, input, change, rote, plug, unplug, cp, quit/exit): ")
		var command string
		_, _ = fmt.Scan(&command)
		switch command {
		case "info":
			var rotorNo []int
			for _, r := range c.enigmaMachine.Rotors {
				rotorNo = append(rotorNo, r.No())
			}
			fmt.Printf("Current rotor no: %v\n", rotorNo)
			fmt.Printf("Current rotor position: %v\n", utils.RuneSliceToStringSlice(c.enigmaMachine.GetRotorPosition()))
			plugs := c.enigmaMachine.PlugBoard.PluggedCables()
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
				r, _ = c.enigmaMachine.TypeKey(machine.KeyboardKey(r))
				res += string(r)
			}
			fmt.Printf("Output: %s\n", res)
			fmt.Printf("Current rotor position: %v\n", utils.RuneSliceToStringSlice(c.enigmaMachine.GetRotorPosition()))
			fmt.Println("OK")
		case "change":
			fmt.Printf("change rotors, please input %d rotor numbers (seperated by space): ", c.conf.RotorCount)
			rotorIndexes := make([]int, c.conf.RotorCount)
			for i := 0; i < c.conf.RotorCount; i++ {
				_, _ = fmt.Scan(&rotorIndexes[i])
			}
			fmt.Printf("Got no: %v\n", rotorIndexes)
			var rotors []*machine.Rotor
			for _, n := range rotorIndexes {
				rot, err := c.conf.MakeRotorByNo(n)
				if err != nil {
					fmt.Printf("invalid rotor (no=%d): %v", n, err)
					continue mainLoop
				}
				_ = rot.SetPosition('A')
				rotors = append(rotors, rot)
			}
			c.enigmaMachine.Rotors = rotors
			fmt.Println("OK")
		case "rote":
			fmt.Printf("rotate rotors, please input %d alphabets (altogether without space): ", c.conf.RotorCount)
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
			err := c.enigmaMachine.SetRotorPosition(rarr)
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
			err := c.enigmaMachine.PlugBoard.Plug(rune(str[0]), rune(str[1]))
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
			err := c.enigmaMachine.PlugBoard.Unplug(rune(str[0]), rune(str[1]))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue mainLoop
			}
			fmt.Println("OK")
		case "cp":
			fmt.Printf("Clearing plugs...")
			c.enigmaMachine.PlugBoard.Clear()
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
