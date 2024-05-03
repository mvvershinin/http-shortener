package config

import (
	"flag"
	"fmt"
	"strings"
)

var ServerProtocol = "http://"
var ServerAddress = "localhost:8080"
var APIPrefix = ""

type Config struct {
	ServerProtocol string
	ServerAddress  string
	APIPrefix      string
}

func (c Config) GetServerLINK() string {
	return fmt.Sprintf("%s%s/%s", c.ServerProtocol, c.ServerAddress, GetAPIPrefixString(c.APIPrefix))
}

func GetConfig() Config {
	flag.StringVar(&ServerAddress, "a", ServerAddress, "The address and port to listen on")
	flag.StringVar(&APIPrefix, "b", APIPrefix, "Api prefix to listen on")
	flag.Parse()

	c := Config{
		ServerProtocol: ServerProtocol,
		ServerAddress:  ServerAddress,
		APIPrefix:      strings.Trim(APIPrefix, "/"),
	}

	return c
}

func GetAPIPrefixString(prefix string) string {
	if len(prefix) > 1 {
		return prefix + "/"
	} else {
		return ""
	}
}
