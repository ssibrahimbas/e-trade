package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	Server struct {
		Host string `env:"SERVER_HOST"`
		Port string `env:"SERVER_PORT"`
	}
	MongoDB struct {
		Uri    string `env:"MONGO_URI"`
		DbName string `env:"MONGO_DB_NAME"`
	}
}

func New() *Config {
	c := &Config{}
	file, err := os.Open("./config.env")
	err = dotenv.NewDecoder(file).Decode(c)
	if err != nil {
		panic(err)
	}
	return c
}
