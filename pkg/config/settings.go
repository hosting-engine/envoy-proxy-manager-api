package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Settings represents app configuration
type Settings struct {
	Debug                bool   `envconfig:"APP_DEBUG" default:"false"`
	Port                 string `envconfig:"APP_PORT" default:"8080"`
	Timeout              uint16 `envconfig:"API_TIMEOUT" default:"600"`
	Version              string `envconfig:"APP_VERSION" default:"dev"`
	DatabaseDSN          string `envconfig:"DATABASE_DSN" default:""`
	DatabaseType         string `envconfig:"DATABASE_TYPE" default:""`
	EnableAdminInterface bool   `envconfig:"ENABLE_ADMIN_INTERFACE" default:""`
	AdminInterfacePath   string `envconfig:"ADMIN_INTERFACE_PATH" default:"./ui"`
}

// NewSettings Provides a new value of settings
func NewSettings(prefix string) Settings {
	s := Settings{}
	if err := envconfig.Process(prefix, s); err != nil {
		log.Fatal(err)
	}
	return s
}
