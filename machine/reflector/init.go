package reflector

import (
	"enigma/config"
	"fmt"
)

var (
	Configuration *Reflector // The system initializes with some predefined reflector
)

func init() {
	if config.Conf.Reflector == nil {
		panic("no reflector config found")
	}
	r, err := NewReflectorFromConf(config.Conf.Reflector)
	if err != nil {
		panic(fmt.Errorf("reflector err: %w", err))
	}
	Configuration = r
}
