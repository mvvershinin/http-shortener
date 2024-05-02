package config

import (
	"fmt"
)

const ServerProtocol = "http://"
const ServerAddress = "localhost:8888"
const APIPrefix = "/api"

type Config struct {
	ServerProtocol string
	ServerAddress  string
	APIPrefix      string
}

func (c Config) GetServerLINK() string {
	return fmt.Sprintf("%s%s%s", c.ServerProtocol, c.ServerAddress, c.APIPrefix)
}

func GetConfig() Config {
	c := Config{
		ServerProtocol: ServerProtocol,
		ServerAddress:  ServerAddress,
		APIPrefix:      APIPrefix,
	}

	return c
}
