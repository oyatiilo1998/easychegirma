package config

import (
	"log"

	"github.com/MrWebUzb/goenv"
)

type Config struct {
	Environment   string `env:"ENVIRONMENT" default:"develop"`
	MongoHost     string `env:"MONGO_HOST" default:"localhost"`
	MongoPort     int    `env:"MONGO_PORT" default:"27017"`
	MongoDatabase string `env:"MONGO_DATABASE" default:"chegirma_setting_service"`
	MongoPassword string `env:"MONGO_PASSWORD" default:"root"`
	MongoUser     string `env:"MONGO_USER" default:"root"`
	LogLevel      string `env:"LOG_LEVEL" default:"debug"`
	RPCPort       string `env:"RPC_PORT" default:":8003"`
}

func Load(fileNames ...string) Config {
	config := Config{}

	env, err := goenv.New(fileNames...)

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	if err := env.Parse(&config, nil); err != nil {
		panic(err)
	}

	return config
}
