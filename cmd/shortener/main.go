package main

import (
	"flag"
	"fmt"
	"github.com/mvvershinin/http-shortener/config"
	"github.com/mvvershinin/http-shortener/internal/app/handler"
	"net/http"
)

var Cfg config.Config
var serverAddress string
var APIPrefix string

func init() {
	flag.StringVar(&serverAddress, "a", Cfg.ServerAddress, "The address and port to listen on")
	flag.StringVar(&APIPrefix, "b", Cfg.APIPrefix, "Api prefix to listen on")
}

func main() {
	flag.Parse()
	Cfg = config.GetConfig()
	if len(serverAddress) > 1 {
		Cfg.ServerAddress = serverAddress
	}

	if len(APIPrefix) > 1 {
		Cfg.APIPrefix = APIPrefix + "/"
	}
	r := handler.GetRouter(Cfg)
	fmt.Printf("listen on %s\n", Cfg.GetServerLINK())
	//fmt.Print(config.GetAPIPrefixString(Cfg.APIPrefix))
	fmt.Println(Cfg.ServerAddress)
	err := http.ListenAndServe(Cfg.ServerAddress, r)
	if err != nil {
		panic(err)
	}
}
