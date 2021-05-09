package util_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/ocsen-hoc-code/go-auj/builder"
	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/models/service"
)

const (
	GET        = "GET"
	POST       = "POST"
	PUT        = "PUT"
	DELETED    = "DELETED"
	NONE_TOKEN = ""
	OLD_JWT    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjA1NTcyNzgsInVzZXJfaWQiOiJlNTJlNDZhZC1jNjU1LTQwMTEtODcxYy1kY2U2MDFhOTg2ZGUiLCJ1c2VyX25hbWUiOiJmaXJzdFVzZXIgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICJ9.6CUqH4Ygqy1cNi0fLpDQ_wQ1Gzqu5O_ESAesVsy8xbg"
)
const USERID = "e52e46ad-c655-4011-871c-dce601a986de"
const USERNAME = "firstUser"
const PASSWORD = "example"

func InstallTest() service.Service {
	serv := service.Service{
		Server:    gin.Default(),
		SecretKey: config.NewJWTConfig(&config.JWTConfig{SecretKey: "Minh dep trai", ExpireTime: 900}),
		Config:    config.NewDbConfig(&config.DbConfig{TestEnv: true})}

	service.NewService(serv)
	container := builder.BuildContainer()
	container.Invoke(func(serv *service.Service) {})

	return serv
}

func CreateRequest(route *gin.Engine, method string, url string, token string, jsonData []byte, params ...map[string]string) (int, map[string]interface{}) {
	var body map[string]interface{}
	var req *http.Request

	res := httptest.NewRecorder()
	if nil != jsonData {
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}

	switch method {
	case GET:

		if nil != params && len(params) > 0 {
			query := req.URL.Query()
			for key, val := range params[0] {
				query.Add(key, val)
			}
			req.URL.RawQuery = query.Encode()
		}
		break
	}
	if "" != token {
		req.Header.Set("Authorization", token)
	}
	route.ServeHTTP(res, req)

	json.NewDecoder(res.Body).Decode(&body)

	return res.Code, body
}
