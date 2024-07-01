package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type otel struct {
	IsSecure bool   `default:"false"`
	URI      string `default:"otel-collector"`
	Port     string `default:"4317"`
}

type service struct {
	Name string `default:"gin-example"`
	Port string `default:"8080"`
}

type config struct {
	Otel    otel
	Service service
}

var Config config

func InitConfig() {
	err := envconfig.Process("", &Config)

	if err != nil {
		log.Fatal(err.Error())
	}
}
