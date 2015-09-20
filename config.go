package shot

import (
	"gopkg.in/gcfg.v1"
)

type Config struct {
	App struct {
		PhantomJsPath string
		Workers       int
		Http          bool
		Port          int

		CachePath string
	}
}

func NewConfig(filename string) (*Config, error) {
	cfg := new(Config)
	e := gcfg.ReadFileInto(cfg, filename)
	return cfg, e
}
