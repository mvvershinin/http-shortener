package handler

import (
	"fmt"
	"github.com/mvvershinin/http-shortener/internal/app/config"
	"github.com/mvvershinin/http-shortener/internal/app/strencoder"
	"io"
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
	str, _ := io.ReadAll(req.Body)
	encoded := strencoder.EncodeStr(string(str))
	link := fmt.Sprintf("%s/%s", cfg.GetServerURL(), encoded)
	res.Header().Add("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	var _, err = res.Write([]byte(fmt.Sprintf("%v", link)))
	if err != nil {
		return
	}
}

func getHandler(res http.ResponseWriter, req *http.Request) {
	//todo get url part decode
	str := cfg.DefaultURL
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
