package handler

import (
	"fmt"
	"github.com/mvvershinin/http-shortener/app/config"
	"github.com/mvvershinin/http-shortener/app/strencoder"
	"net/http"
)

func MainHandler(res http.ResponseWriter, req *http.Request) {
	cfg := config.GetConfig()
	if http.MethodGet != req.Method && req.Method != http.MethodPost {
		res.WriteHeader(http.StatusBadRequest)
		errorMessage := "Bad Request: Only requests to /link/ are allowed."
		_, err := res.Write([]byte(errorMessage))
		if err != nil {
			return
		}
		return
	}
	if req.Method == http.MethodPost {
		str := strencoder.EncodeStr(cfg.DefaultUrl)
		res.Header().Add("content-type", "text/plain")
		var _, err = res.Write([]byte(fmt.Sprintf("%v", str)))
		if err != nil {
			return
		}
		return
	}
	if req.Method == http.MethodGet {
		res.Header().Add("content-type", "text/plain")
		_, err := res.Write([]byte(cfg.DefaultUrl))
		if err != nil {
			return
		}
		return
	}
}
