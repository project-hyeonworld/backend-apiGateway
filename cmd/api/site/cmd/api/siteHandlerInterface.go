package site

import (
	"way-manager/api/shared/common"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	common.IHandler
	Add(c *gin.Context)
}