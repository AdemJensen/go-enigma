package main

import (
	"encoding/json"
	"enigma/config"
	"enigma/utils"
	"flag"
	"io/ioutil"
)

func main() {
	var (
		rotorCount  int
		nGenRotors  int
		outFileName string
	)

	flag.IntVar(&rotorCount, "rotor_count", 3, "Defines the rotor count of this machine")
	flag.IntVar(&nGenRotors, "num_gen_rotors", 5, "How many rotors to generate")
	flag.StringVar(&outFileName, "out_file_name", "conf.json", "Output config to which file")
	flag.Parse()

	// rotors
	var rotors []*config.RotorConfig
	for i := 0; i < nGenRotors; i++ {
		rotors = append(rotors, &config.RotorConfig{
			No:      i + 1,
			Mapping: utils.RuneMapToStringMap(utils.RandomAlphabetMappingTable()),
		})
	}

	// initial rotors
	var initialRotors []int
	for i := 0; i < rotorCount; i++ {
		initialRotors = append(initialRotors, i+1)
	}

	// reflector
	var allAlphabets []rune
	for key := 'A'; key <= 'Z'; key++ {
		allAlphabets = append(allAlphabets, key)
	}
	mappingKeys := utils.RandomSelectRuneArray(allAlphabets, 13)
	val := utils.DeleteFromRuneSlice(allAlphabets, mappingKeys...)
	mappingTable := make(map[rune]rune)
	for idx, key := range mappingKeys {
		mappingTable[key] = val[idx]
		mappingTable[val[idx]] = key
	}

	conf := &config.Config{
		Debug:         false,
		RotorCount:    rotorCount,
		InitialRotors: initialRotors,
		InitialPlugs:  []config.PlugConfig{},
		Rotors:        rotors,
		Reflector:     utils.RuneMapToStringMap(mappingTable),
	}

	str, _ := json.MarshalIndent(conf, "", "    ")
	err := ioutil.WriteFile(outFileName, str, 0644) // the 0644 is octal representation of the file mode
	if err != nil {
		panic("failed to write file: " + err.Error())
	}
}
