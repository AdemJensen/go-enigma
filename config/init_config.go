package config

import (
	"encoding/json"
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

type PlugConfig []string

type RotorConfig struct {
	No      int               `json:"no"`
	Mapping map[string]string `json:"mapping"`
}

type ReflectorConf map[string]string

var Conf Config

func init() {
	f, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		panic("failed to read file conf.json: " + err.Error())
	}
	err = json.Unmarshal(f, &Conf)
	if err != nil {
		panic("failed to unmarshal conf: " + err.Error())
	}
}
