package api

const (
	defaultHostname = "localhost"
	defaultPort     = 8080
)

// Config contains settings of the client.
type Config struct {
	Hostname string `mapstructure:"hostname" json:"hostname,omitempty"`
	Port     int    `mapstructure:"port" json:"port,omitempty"`
	Swagger  bool   `mapstructure:"swagger" json:"swagger,omitempty"`
}

// NewConfig returns a new Config instance.
func NewConfig() *Config {
	return &Config{
		Hostname: defaultHostname,
		Port:     defaultPort,
	}
}
