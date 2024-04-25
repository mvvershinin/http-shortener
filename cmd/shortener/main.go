package main

import (
	"fmt"
	"http-shortener/app/handler"
	"http-shortener/internal/app/config"
	"net/http"
)

func main() {
	cfg := config.GetConfig()
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handler.MainHandler)
	fmt.Println(fmt.Sprintf("listen %s", cfg.GetServerUrl()))
	err := http.ListenAndServe(cfg.GetServerUrl(), mux)
	if err != nil {
		panic(err)
	}
}
