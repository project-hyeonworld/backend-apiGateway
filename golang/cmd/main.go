package main

import (
	nginxLib "way-manager/api/nginx/cmd/lib"

	siteLib "way-manager/api/site/cmd/lib"
	secret "way-manager/configuration/secret"

	"github.com/gin-gonic/gin"
)

func main() {
	secretValue := secret.Value{}
	secretValue.Init()
	router := gin.Default()

	startApiServer(secretValue, router)
	startGrpcServer(secretValue)
	router.Run(":5000")
}

func startGrpcServer(secretValue secret.Value) {
	nginxLib.StartGrpcServer(&secretValue)
}

func startApiServer(secretValue secret.Value, router *gin.Engine) {
	siteLib.StartApiServer(&secretValue, router)
	nginxLib.StartApiServer(&secretValue, router)
}
