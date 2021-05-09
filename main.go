package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ocsen-hoc-code/go-auj/builder"
	"github.com/ocsen-hoc-code/go-auj/models/config"
	"github.com/ocsen-hoc-code/go-auj/models/service"
)

func SetupServer(serv *service.Service) *service.Service {
	service.NewService(serv)
	container := builder.BuildContainer()
	container.Invoke(func(serv *service.Service) {})
	return serv
}

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("Can't loading .env file")
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if nil != err {
		port = 5432
	}
	expireTime, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRED_TIME"))
	if nil != err {
		expireTime = 900
	}
	serv := SetupServer(
		&service.Service{
			Server:    gin.Default(),
			Port:      8888,
			SecretKey: config.NewJWTConfig(&config.JWTConfig{SecretKey: "Minh dep trai", ExpireTime: expireTime}),
			Config: config.NewDbConfig(&config.DbConfig{
				Hostname:   os.Getenv("DB_HOST"),
				UserName:   os.Getenv("DB_USERNAME"),
				Password:   os.Getenv("DB_PASSWORD"),
				Port:       port,
				DbName:     os.Getenv("DB_NAME"),
				Drivername: os.Getenv("DB_DRIVER"),
			})})
	serv.Run()
}
