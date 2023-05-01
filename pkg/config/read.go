package config

import (
	"fmt"
	"log"
	"os"

	"deall-alfon/pkg/util/fn"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func readJSONConfig(cfg *ConfigStr, basePath, cfgLocation string, fileList []string) {
	op := fn.Name()

	v := viper.New()
	v.AddConfigPath(fmt.Sprintf("%s/%s/", basePath, cfgLocation))
	for _, file := range fileList {
		v.SetConfigName(file)
		err := v.MergeInConfig()
		if err != nil {
			log.Fatalf("[%v] fail merge config for file %v, error: %v", op, file, err)
		}
	}

	err := v.Unmarshal(cfg, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "json"
	})
	if err != nil {
		log.Fatalf("[%v] fail unmarshal config, error: %v", op, err)
	}
}

func readEnvironment(cfg *ConfigStr) {
	if env := os.Getenv(Environment); env != "" {
		cfg.Environment = env
	}
}
