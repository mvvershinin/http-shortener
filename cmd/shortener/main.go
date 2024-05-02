package main

import (
	"flag"
	"fmt"
	"github.com/mvvershinin/http-shortener/config"
	"github.com/mvvershinin/http-shortener/internal/app/handler"
	"net/http"
)

var Cfg = config.GetConfig()

func init() {
	flag.StringVar(&Cfg.ServerAddress, "a", Cfg.ServerAddress, "The address and port to listen on")
	flag.StringVar(&Cfg.APIPrefix, "b", Cfg.APIPrefix, "Api prefix to listen on")
}

func main() {
	flag.Parse()
	r := handler.GetRouter(Cfg)
	fmt.Printf("listen on %s\n", Cfg.GetServerLINK())
	err := http.ListenAndServe(Cfg.ServerAddress, r)
	if err != nil {
		panic(err)
	}
}
