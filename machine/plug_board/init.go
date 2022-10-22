package plug_board

import (
	"enigma/config"
	"fmt"
)

var (
	Configuration *PlugBoard
)

func init() {
	Configuration = NewPlugBoard()
	err := Configuration.ApplyConfig(config.Conf.InitialPlugs)
	if err != nil {
		panic(fmt.Errorf("plug config err: %v", err))
	}
}
