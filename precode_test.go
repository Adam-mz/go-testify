package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req, _ := http.NewRequest("GET", "/cafe?city=moscow&count="+strconv.Itoa(totalCount+1), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидается статус 200 OK")

	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, totalCount, len(cafes), "Ожидается, что вернутся все доступные кафе")
}

func TestMainHandlerWithValidRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидается статус 200 OK")

	assert.NotEmpty(t, responseRecorder.Body.String(), "Ожидается, что ответ не будет пустым")
}

func TestMainHandlerWithUnsupportedCity(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cafe?city=unknown&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидается статус 400 Bad Request")

	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Ожидается сообщение об ошибке 'wrong city value'")
}
