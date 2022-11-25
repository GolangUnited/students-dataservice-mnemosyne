package configs

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
)

const configPath = "configs/config.yml"

type Config struct {
	RestPort int `yaml:"restPort"`
	GrpcPort int `yaml:"grpcPort"`
}

// Init config from yml file
func Init() (*Config, error) {
	var cfg Config

	rawYAML, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "parsing yaml")
	}

	return &cfg, nil
}
