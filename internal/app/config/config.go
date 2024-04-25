package config

import (
	"fmt"
	"strconv"
)

const serverAddress = "localhost"
const port = 8080
const defaultLink = "https://practicum.yandex.ru"

type Config struct {
	ServerAddress string
	ServerPort    string
	DefaultUrl    string
}

func (c Config) GetServerUrl() string {
	return fmt.Sprintf("%s:%s", c.ServerAddress, c.ServerPort)
}

func GetConfig() Config {
	c := Config{
		ServerAddress: serverAddress,
		ServerPort:    strconv.Itoa(port),
		DefaultUrl:    defaultLink,
	}

	return c
}
