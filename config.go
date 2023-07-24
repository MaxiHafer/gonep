package gonep

import "github.com/kelseyhightower/envconfig"

const ConfigPrefix = "NEP_VIEWER"

type Config struct {
	BaseURL string `default:"user.nepviewer.com"`
	Scheme  string `default:"https"`

	User     string
	Password string
}

func (c Config) FromEnv() error {
	return envconfig.Process(ConfigPrefix, &c)
}
