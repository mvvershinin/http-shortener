package config

import (
	"fmt"
)

const ServerProtocol = "http://"
const ServerAddress = "localhost:8888"
const ApiPrefix = "/api"

type Config struct {
	ServerProtocol string
	ServerAddress  string
	ApiPrefix      string
}

func (c Config) GetServerLINK() string {
	return fmt.Sprintf("%s%s%s", c.ServerProtocol, c.ServerAddress, c.ApiPrefix)
}

func GetConfig() Config {
	c := Config{
		ServerProtocol: ServerProtocol,
		ServerAddress:  ServerAddress,
		ApiPrefix:      ApiPrefix,
	}

	return c
}
