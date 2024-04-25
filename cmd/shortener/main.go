package main

import (
	"fmt"
	"github.com/mvvershinin/http-shortener/internal/app/strencoder"
	"net/http"
)

const serverAddress = "localhost"
const port = 8088
const defaultLink = "https://practicum.yandex.ru/gfgdfgd44gdfdf"

func linkHandler(res http.ResponseWriter, req *http.Request) {
	if http.MethodGet != req.Method && req.Method != http.MethodPost {
		http.Error(res, "Bad Request", http.StatusBadRequest)
		return
	}
	if req.Method == http.MethodPost {
		res.Header().Add("content-type", "text/plain")
		str := strencoder.EncodeStr("https://practicum.yandex.ru")
		fmt.Println(str)
		//		result := getServerUrl() + "/gfSVNBsFs"
		var _, err = res.Write([]byte(fmt.Sprintf("%v", str)))
		if err != nil {
			return
		}
	}
	if req.Method == http.MethodGet {
		//path := req.URL.Path
		res.Header().Add("content-type", "text/plain")
		_, err := res.Write([]byte(defaultLink))
		if err != nil {
			return
		}
	}

	//
}

func getServerUrl() string {
	return fmt.Sprintf("%s:%d", serverAddress, port)
}

func main() {
	server := getServerUrl()
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, linkHandler)

	fmt.Println(fmt.Sprintf("listen %s", server))
	err := http.ListenAndServe(server, mux)
	if err != nil {
		panic(err)
	}

}
