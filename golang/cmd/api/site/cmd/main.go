package main

import (
	"fmt"
	siteLib "way-manager/api/site/cmd/lib"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	if err := siteLib.Start(router); err != nil {
		fmt.Errorf("Failed to start site api: %v", err)
	}

	if err := router.Run(":5000"); err != nil {
		fmt.Errorf("Failed to run server: %v", err)
	}
}
