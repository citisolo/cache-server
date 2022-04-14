package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JSainsburyPLC/third-party-token-server/app"
	"github.com/JSainsburyPLC/third-party-token-server/app/handler"
	"github.com/JSainsburyPLC/third-party-token-server/config"
	"github.com/stretchr/testify/assert"
)


func TestInsertObj(t *testing.T) {
	app := newApp()
	
	var input = []byte(`{"id":"1234", "token":"56789101112131415"}`)
	postDataRequest, _ := http.NewRequest("POST", "/token", bytes.NewBuffer(input))

	postRequestHandler := http.HandlerFunc(app.NewHandler(handler.PostToken))
	postResponse := httptest.NewRecorder()

	postRequestHandler.ServeHTTP(postResponse, postDataRequest)

	status := postResponse.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	getDataRequest, _ := http.NewRequest("GET", "/token?id=1234", nil) 

	getResponse := httptest.NewRecorder()
	getRequestHandler := http.HandlerFunc(app.NewHandler(handler.GetToken))

	getRequestHandler.ServeHTTP(getResponse, getDataRequest)

	status = getResponse.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	var data handler.TokenResponse
	json.NewDecoder(getResponse.Body).Decode(&data)

	//Assert
	assert.NotNil(t, data.Token)
	assert.Equal(t, "56789101112131415",  data.Token)
}

func newApp() *app.App {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	return app
}
