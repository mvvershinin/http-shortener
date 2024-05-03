package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/mvvershinin/http-shortener/config"
	"github.com/mvvershinin/http-shortener/internal/app/strencoder"
	"io"
	"log"
	"net/http"
	"path"
)

const errorMessage = "Bad Request: Something wrong happened."

var Cfg config.Config

func GetRouter(cfg config.Config) *chi.Mux {
	Cfg = cfg
	router := chi.NewRouter()
	router.Route(Cfg.APIPrefix, func(router chi.Router) {
		router.Get("/{uid}", GetHandler)
		router.Post("/", PostHandler)
	})
	router.NotFound(BadRequestHandler)
	router.MethodNotAllowed(BadRequestHandler)

	return router
}

func BadRequestHandler(res http.ResponseWriter, r *http.Request) {
	res.WriteHeader(http.StatusBadRequest)
	fmt.Println(r)
	_, err := res.Write([]byte(errorMessage))
	if err != nil {
		return
	}
}

func PostHandler(res http.ResponseWriter, req *http.Request) {
	str, errRead := io.ReadAll(req.Body)
	if errRead != nil {
		log.Fatal(errRead)
		return
	}
	encoded := strencoder.Base64Encode(string(str))
	link := Cfg.ServerProtocol + path.Join(Cfg.GetServerPath(), encoded)
	res.Header().Add("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	var _, errWrite = res.Write([]byte(link))
	if errWrite != nil {
		log.Fatal(errWrite)
		return
	}
}

func GetHandler(res http.ResponseWriter, req *http.Request) {
	uid := chi.URLParam(req, "uid")
	str, err := strencoder.Base64Decode(uid)
	if err != nil {
		BadRequestHandler(res, req)
	}
	res.Header().Add("content-type", "text/plain")
	res.Header().Add("Location", string(str))
	res.WriteHeader(http.StatusTemporaryRedirect)
}
