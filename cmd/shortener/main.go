package main

import (
	"fmt"
	"github.com/mvvershinin/http-shortener/config"
	"github.com/mvvershinin/http-shortener/internal/app/handler"
	"net/http"
)

var Cfg config.Config

func main() {
	Cfg = config.GetConfig()
	r := handler.GetRouter(Cfg)
	fmt.Printf("listen on %s\n", Cfg.GetServerLINK())
	err := http.ListenAndServe(Cfg.ServerAddress, r)
	if err != nil {
		panic(err)
	}
}
