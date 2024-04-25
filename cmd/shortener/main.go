package main

import (
	"github.com/mvvershinin/http-shortener/internal/app/config"
	"github.com/mvvershinin/http-shortener/internal/app/handler"
	"net/http"
)

func main() {
	cfg := config.GetConfig()
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handler.MainHandler)
	//fmt.Printf("listen %s\r\n", cfg.GetServerURL())
	err := http.ListenAndServe(cfg.GetServerURL(), mux)
	if err != nil {
		panic(err)
	}
}
