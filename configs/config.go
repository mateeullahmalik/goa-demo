package configs

import (
	"encoding/json"

	"github.com/mateeullahmalik/goa-demo/api"
)

const (
	defaultLogLevel = "info"
)

// Config contains configuration of all components
type Config struct {
	LogLevel string      `mapstructure:"log-level" json:"log-level,omitempty"`
	LogFile  string      `mapstructure:"log-file" json:"log-file,omitempty"`
	Quiet    bool        `mapstructure:"quiet" json:"quiet"`
	TempDir  string      `mapstructure:"temp-dir" json:"temp-dir"`
	WorkDir  string      `mapstructure:"work-dir" json:"work-dir"`
	API      *api.Config `mapstructure:"api" json:"api,omitempty"`
}

func (config *Config) String() string {
	// The main purpose of using a custom converting is to avoid unveiling credentials.
	// All credentials fields must be tagged `json:"-"`.
	data, _ := json.Marshal(config)
	return string(data)
}

// New returns a new Config instance
func New() *Config {
	return &Config{
		LogLevel: defaultLogLevel,
		API:      api.NewConfig(),
	}
}
