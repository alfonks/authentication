package config

import (
	"log"
	"path/filepath"
	"sync"

	"deall-alfon/pkg/util/fn"
)

var (
	configOnce sync.Once
	config     ConfigStr
)

func GetConfig() ConfigStr {
	configOnce.Do(func() {
		op := fn.Name()

		readEnvironment(&config)
		configListJSON := []string{
			"deall-alfon-config",
			"deall-alfon-secret-config",
		}
		path, err := filepath.Abs(".")
		if err != nil {
			log.Fatalf("[%v] cant find absolute path: %v", op, err)
		}
		configLocation := "cfg"
		readJSONConfig(&config, path, configLocation, configListJSON)
	})
	return config
}
