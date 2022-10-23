package json_conf

import (
	"encoding/json"
	"enigma/machine"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Debug         bool           `json:"debug"`
	RotorCount    int            `json:"rotor_count"`
	InitialRotors []int          `json:"initial_rotors"`
	InitialPlugs  []PlugConfig   `json:"initial_plugs"`
	Rotors        []*RotorConfig `json:"rotors"`
	Reflector     ReflectorConf  `json:"reflector"`
}

func (c *Config) MakeInitialRotors() ([]*machine.Rotor, error) {
	if c.RotorCount != len(c.InitialRotors) {
		return nil, fmt.Errorf("amount of initial rotors must equals to rotor count")
	}
	var initialRotors []*machine.Rotor
	for _, rNo := range c.InitialRotors {
		rot, err := c.MakeRotorByNo(rNo)
		if err != nil {
			return nil, fmt.Errorf("make rotor %d failed: %w", rNo, err)
		}
		initialRotors = append(initialRotors, rot)
	}
	return initialRotors, nil
}

func (c *Config) MakeRotorByNo(no int) (*machine.Rotor, error) {
	for _, cnf := range c.Rotors {
		if cnf == nil || cnf.No != no {
			continue
		}
		rot, err := cnf.MakeRotor()
		if err != nil {
			return nil, fmt.Errorf("make rotor %d failed: %w", cnf.No, err)
		}
		return rot, nil
	}
	return nil, fmt.Errorf("rotor no %d not exist", no)
}

func (c *Config) MakeInitialPlugBoard() (*machine.PlugBoard, error) {
	plg := machine.NewPlugBoard()
	for idx, cnf := range c.InitialPlugs {
		if cnf == nil {
			continue
		}
		err := cnf.ApplyTo(plg)
		if err != nil {
			return nil, fmt.Errorf("error when apply index %d: %w", idx, err)
		}
	}
	return plg, nil
}

func NewConfigFromFile(filename string) (*Config, error) {
	var conf Config
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	err = json.Unmarshal(f, &conf)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal conf: %w", err)
	}
	return &conf, nil
}
