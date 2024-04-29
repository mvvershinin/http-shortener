package main

import (
	"github.com/mvvershinin/http-shortener/internal/app/handler"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerGetSuccess(t *testing.T) {
	contentType := "text/plain"
	testCases := []struct {
		method                 string
		path                   string
		expectedCode           int
		expectedContentType    string
		expectedHeaderLocation string
	}{
		{
			method:                 http.MethodGet,
			path:                   "/aHR0cHM6Ly95YW5kZXgucnUv",
			expectedCode:           http.StatusTemporaryRedirect,
			expectedContentType:    contentType,
			expectedHeaderLocation: "https://yandex.ru/",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/"+tc.path, nil)
			w := httptest.NewRecorder()

			handler.MainHandler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tc.expectedContentType, w.Header().Get("Content-Type"))
			assert.Equal(t, tc.expectedHeaderLocation, w.Header().Get("Location"))
		})
	}
}

func TestHandlerPostSuccess(t *testing.T) {
	successBody := "http://localhost:8080/aHR0cHM6Ly95YW5kZXgucnUv"
	requestBody := "https://yandex.ru/"
	contentType := "text/plain"
	testCases := []struct {
		method              string
		requestBody         string
		expectedCode        int
		expectedContentType string
		expectedBody        string
	}{
		{
			method:              http.MethodPost,
			requestBody:         requestBody,
			expectedCode:        http.StatusCreated,
			expectedContentType: contentType,
			expectedBody:        successBody,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", strings.NewReader(tc.requestBody))
			w := httptest.NewRecorder()

			handler.MainHandler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tc.expectedContentType, w.Header().Get("Content-Type"))
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}

func TestErrorsWebhook(t *testing.T) {
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
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", nil)
			w := httptest.NewRecorder()

			handler.MainHandler(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
		})
	}
}
