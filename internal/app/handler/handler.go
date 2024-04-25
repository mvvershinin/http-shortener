package handler

import (
	"fmt"
	"http-shortener/internal/app/config"
	"http-shortener/internal/app/strencoder"
	"net/http"
)

const errorMessage = "Bad Request: Only requests GEt and POST are allowed."

var cfg = config.GetConfig()

func badRequestHandler(res http.ResponseWriter) {
	res.WriteHeader(http.StatusBadRequest)

	_, err := res.Write([]byte(errorMessage))
	if err != nil {
		return
	}
}

func postHandler(res http.ResponseWriter, req *http.Request) {
	//todo post body encode
	str := strencoder.EncodeStr(cfg.DefaultUrl)
	res.Header().Add("content-type", "text/plain")
	var _, err = res.Write([]byte(fmt.Sprintf("%v", str)))
	if err != nil {
		return
	}
}

func getHandler(res http.ResponseWriter, req *http.Request) {
	//todo get url part decode
	str := cfg.DefaultUrl
	res.Header().Add("content-type", "text/plain")
	res.Header().Add("Location", str)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func MainHandler(res http.ResponseWriter, req *http.Request) {

	if http.MethodGet != req.Method && req.Method != http.MethodPost {
		badRequestHandler(res)

		return
	}
	if req.Method == http.MethodPost {
		postHandler(res, req)

		return
	}
	if req.Method == http.MethodGet {
		getHandler(res, req)

		return
	}
}
