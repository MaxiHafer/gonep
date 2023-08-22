package gonep

import (
	"github.com/kelseyhightower/envconfig"
)

const ConfigPrefix = "NEP_VIEWER"

type Config struct {
	User     string `required:"true"`
	Password string `required:"true"`
}

func (c *Config) FromEnv() error {
	return envconfig.Process(ConfigPrefix, c)
}
