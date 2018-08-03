package config

import "github.com/JREAMLU/j-kit/go-micro/util"

const (
	name    = "example"
	version = "v1"
)

// HelloConfig hello config
type HelloConfig struct {
	*util.Config

	Hello struct {
		Secret string
	}
}

// Load load config
func Load() (*HelloConfig, error) {
	// load redis mysql elastic client

	// load parent config
	config := &HelloConfig{}
	err := util.LoadCustomConfig("10.200.202.35:8500", name, version, config)

	return config, err
}
