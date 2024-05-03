package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"log"
	"strings"
)

var ServerProtocol = "http://"
var ServerAddress = "localhost:8080"
var APIPrefix = ""

type Config struct {
	ServerProtocol string
	ServerAddress  string `env:"SERVER_ADDRESS"`
	APIPrefix      string `env:"BASE_URL"`
}

func (c Config) GetServerPath() string {
	return fmt.Sprintf("%s%s", c.ServerAddress, c.APIPrefix)
}

func (c Config) GetServerLINK() string {
	return fmt.Sprintf("%s%s%s", c.ServerProtocol, c.ServerAddress, c.APIPrefix)
}

func GetConfig() Config {
	cfg := Config{
		ServerProtocol: ServerProtocol,
		ServerAddress:  ServerAddress,
		APIPrefix:      APIPrefix,
	}

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "The address and port to listen on")
	flag.StringVar(&cfg.APIPrefix, "b", cfg.APIPrefix, "Api prefix to listen on")
	flag.Parse()

	c := Config{
		ServerProtocol: cfg.ServerProtocol,
		ServerAddress:  cfg.ServerAddress,
		APIPrefix:      "/" + strings.Trim(cfg.APIPrefix, "/"),
	}

	return c
}
