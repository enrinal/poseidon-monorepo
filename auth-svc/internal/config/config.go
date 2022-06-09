package config

import (
	"encoding/json"

	"github.com/enrinal/poseidon-monorepo/auth-svc/internal/log"
	"github.com/spf13/viper"
)

var viperInstance = viper.New()
var Default Config

func init() {
	log.Debug("INIT CONFIG")
}

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		URL string
	}
}

func (d Config) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

// Parse get all config support in app
func Parse() Config {
	if err := viperInstance.Unmarshal(&Default); err != nil {
		log.Fatal("Fail to read configuration", err)
	}
	return Default
}

// Viper instance
func Viper() *viper.Viper {
	return viperInstance
}
