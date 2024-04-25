package config

import (
	"fmt"
	"strconv"
)

const serverProtocol = "http://"
const serverAddress = "localhost"
const port = 8080
const defaultLink = "https://practicum.yandex.ru"

type Config struct {
	ServerProtocol string
	ServerAddress  string
	ServerPort     string
	DefaultURL     string
}

func (c Config) GetServerURL() string {
	return fmt.Sprintf("%s:%s", c.ServerAddress, c.ServerPort)
}

func (c Config) GetServerLINK() string {
	return fmt.Sprintf("%s%s:%s", c.ServerProtocol, c.ServerAddress, c.ServerPort)
}

func GetConfig() Config {
	c := Config{
		ServerProtocol: serverProtocol,
		ServerAddress:  serverAddress,
		ServerPort:     strconv.Itoa(port),
		DefaultURL:     defaultLink,
	}

	return c
}
