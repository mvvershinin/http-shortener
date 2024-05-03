package main

import (
	"context"
	"encoding/base64"
	"github.com/go-chi/chi/v5"
	"github.com/mvvershinin/http-shortener/config"
	"github.com/mvvershinin/http-shortener/internal/app/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerGetSuccess(t *testing.T) {
	contentType := "text/plain"

	cfg := config.Config{
		ServerProtocol: "http://",
		ServerAddress:  "localhost:8080",
		APIPrefix:      "/",
	}
	router := handler.GetRouter(cfg)
	ts := httptest.NewServer(router)
	defer ts.Close()

	testCases := []struct {
		name                   string
		method                 string
		path                   string
		expectedCode           int
		expectedContentType    string
		expectedHeaderLocation string
	}{
		{
			name:                   "success GET redirect link",
			method:                 http.MethodGet,
			path:                   "aHR0cHM6Ly95YW5kZXgucnUv",
			expectedCode:           http.StatusTemporaryRedirect,
			expectedContentType:    contentType,
			expectedHeaderLocation: "https://yandex.ru/",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", nil)
			w := httptest.NewRecorder()
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("uid", tc.path)

			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

			handler.GetHandler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tc.expectedContentType, w.Header().Get("Content-Type"), "Content-Type не совпадает с ожидаемым")
			assert.Equal(t, tc.expectedHeaderLocation, w.Header().Get("Location"), "Location не совпадает с ожидаемым")
		})
	}
}

func TestHandlerPostSuccess(t *testing.T) {
	cfg := config.Config{
		ServerProtocol: "http://",
		ServerAddress:  "localhost:8080",
		APIPrefix:      "/",
	}
	router := handler.GetRouter(cfg)
	ts := httptest.NewServer(router)
	defer ts.Close()

	requestBody := "https://yandex.ru/"
	str := base64.StdEncoding.EncodeToString([]byte(requestBody))
	successBody := cfg.GetServerLINK() + str
	contentType := "text/plain"
	testCases := []struct {
		name                string
		method              string
		requestBody         string
		expectedCode        int
		expectedContentType string
		expectedBody        string
	}{
		{
			name:                "success POST get redirect link",
			method:              http.MethodPost,
			requestBody:         requestBody,
			expectedCode:        http.StatusCreated,
			expectedContentType: contentType,
			expectedBody:        successBody,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", strings.NewReader(tc.requestBody))
			w := httptest.NewRecorder()

			handler.PostHandler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tc.expectedContentType, w.Header().Get("Content-Type"))
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}

func TestErrors(t *testing.T) {
	testCases := []struct {
		method              string
		path                string
		expectedCode        int
		expectedContentType string
		expectedBody        string
		requestBody         string
	}{
		{method: http.MethodPut, expectedCode: http.StatusBadRequest, expectedBody: ""},
		{method: http.MethodPatch, expectedCode: http.StatusBadRequest, expectedBody: ""},
		{method: http.MethodDelete, expectedCode: http.StatusBadRequest, expectedBody: ""},
	}

	for _, tc := range testCases {
		t.Run("method not allowed "+tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", nil)
			w := httptest.NewRecorder()

			handler.BadRequestHandler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
		})
	}
}
