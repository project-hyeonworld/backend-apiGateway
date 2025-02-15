package main

import (
	"fmt"
	nginxLib "way-manager/api/nginx/cmd/lib"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	if err := nginxLib.StartApiServer(nil, router); err != nil {
		fmt.Errorf("Failed to start site api: %v", err)
		return
	}

	if err := router.Run(":5000"); err != nil {
		fmt.Errorf("Failed to run server: %v", err)
		return
	}
}
