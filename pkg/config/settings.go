package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var S *Setting

// Setting represents app configuration
type Setting struct {
	Port         string `envconfig:"APP_PORT" default:"8080"`
	Timeout      uint16 `envconfig:"API_TIMEOUT" default:"600"`
	Version      string `envconfig:"APP_VERSION" default:"dev"`
	DatabaseDSN  string `envconfig:"DATABASE_DSN" default:""`
	DatabaseType string `envconfig:"DATABASE_TYPE" default:""`
}

func init() {
	S = &Setting{}
	if err := envconfig.Process("", S); err != nil {
		log.Fatal(err)
	}
}
