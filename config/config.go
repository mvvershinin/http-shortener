package config

import (
	"fmt"
	"strings"
)

const ServerProtocol = "http://"
const ServerAddress = "localhost:8080"
const APIPrefix = ""

type Config struct {
	ServerProtocol string
	ServerAddress  string
	APIPrefix      string
}

func (c Config) GetServerLINK() string {
	return fmt.Sprintf("%s%s/%s", c.ServerProtocol, c.ServerAddress, GetAPIPrefixString(c.APIPrefix))
}

func GetConfig() Config {
	c := Config{
		ServerProtocol: ServerProtocol,
		ServerAddress:  ServerAddress,
		APIPrefix:      APIPrefix,
	}

	return c
}

func GetAPIPrefixString(prefix string) string {
	if len(prefix) > 1 {
		return strings.Trim(prefix, "/")
	} else {
		return ""
	}
}
