package handler

import (
	"fmt"
	"http-shortener/internal/app/config"
	"http-shortener/internal/app/strencoder"
	"net/http"
)

func MainHandler(res http.ResponseWriter, req *http.Request) {
	cfg := config.GetConfig()
	if http.MethodGet != req.Method && req.Method != http.MethodPost {
		res.WriteHeader(http.StatusBadRequest)
		errorMessage := "Bad Request: Only requests GEt and POST are allowed."
		_, err := res.Write([]byte(errorMessage))
		if err != nil {
			return
		}
		return
	}
	if req.Method == http.MethodPost {
		//todo post body encode
		str := strencoder.EncodeStr(cfg.DefaultUrl)
		res.Header().Add("content-type", "text/plain")
		var _, err = res.Write([]byte(fmt.Sprintf("%v", str)))
		if err != nil {
			return
		}
		return
	}
	if req.Method == http.MethodGet {
		//todo get url part decode
		res.Header().Add("content-type", "text/plain")
		_, err := res.Write([]byte(cfg.DefaultUrl))
		if err != nil {
			return
		}
		return
	}
}
