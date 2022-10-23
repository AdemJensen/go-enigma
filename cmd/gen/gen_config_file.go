package main

import (
	"encoding/json"
	"enigma/configurator/json_conf"
	"flag"
	"io/ioutil"
)

func main() {
	var (
		rotorCount  int
		nGenRotors  int
		outFileName string
	)

	flag.IntVar(&rotorCount, "rotor_count", 3, "Defines the initial rotor count of this machine")
	flag.IntVar(&nGenRotors, "num_gen_rotors", 5, "How many rotors to generate")
	flag.StringVar(&outFileName, "out_file_name", "conf.json", "Output config to which file")
	flag.Parse()

	if nGenRotors < rotorCount {
		panic("num_gen_rotors must be greater than rotor_count")
	}

	// rotors
	var rotors []*json_conf.RotorConfig
	for i := 0; i < nGenRotors; i++ {
		rotors = append(rotors, json_conf.NewRandomRotorConfig(i+1))
	}

	// initial rotors
	var initialRotors []int
	for i := 0; i < rotorCount; i++ {
		initialRotors = append(initialRotors, i+1)
	}

	// reflector
	refec := json_conf.NewRandomReflectorConfig()

	conf := &json_conf.Config{
		InitialRotors: initialRotors,
		InitialPlugs:  []json_conf.PlugConfig{},
		Rotors:        rotors,
		Reflector:     refec,
	}

	str, _ := json.MarshalIndent(conf, "", "    ")
	err := ioutil.WriteFile(outFileName, str, 0644) // the 0644 is octal representation of the file mode
	if err != nil {
		panic("failed to write file: " + err.Error())
	}
}
