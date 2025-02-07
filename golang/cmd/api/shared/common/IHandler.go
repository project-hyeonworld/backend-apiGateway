package common

import "github.com/gin-gonic/gin"

type IHandler interface {
	Add(c *gin.Context)
}
