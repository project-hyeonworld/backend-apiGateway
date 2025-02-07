package main

import (
	siteLib "way-manager/api/site/cmd/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	siteLib.Start(router)
	router.Run(":5000")
}
