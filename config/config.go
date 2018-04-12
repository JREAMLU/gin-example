package config

const (
	name    = "example"
	version = "v1"
)

// HelloConfig hello config
type HelloConfig struct {
}

// LoadConfig load config
func LoadConfig() (*HelloConfig, error) {
	// loadConfig

	config := &HelloConfig{}
	return config, nil
}

// reload watch reload consul
func reload() {
	// consul config has changed
}
