package main

import (
	"github.com/mvvershinin/http-shortener/config"
	"github.com/mvvershinin/http-shortener/internal/app/router"
	"net/http"
)

func main() {
	cfg := config.GetConfig()
	r := router.GetRouter()
	err := http.ListenAndServe(cfg.GetServerURL(), r)
	if err != nil {
		panic(err)
	}
}
