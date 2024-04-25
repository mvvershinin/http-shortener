package main

import (
	"fmt"
	"github.com/mvvershinin/http-shortener/internal/app/config"
	"github.com/mvvershinin/http-shortener/internal/app/handler"
	"net/http"
)

func main() {
	cfg := config.GetConfig()
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handler.MainHandler)
	fmt.Printf("listen %s\r\n", cfg.GetServerUrl())
	err := http.ListenAndServe(cfg.GetServerUrl(), mux)
	if err != nil {
		panic(err)
	}
}
