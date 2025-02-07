package site

import (
	"net/http"
	common "way-manager/api/shared/common"
	model "way-manager/api/shared/common/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	ctrl common.IController
}

func NewHandler(ctrl common.IController) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) Add(c *gin.Context) {
	var input model.ProxyServer
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := h.ctrl.Add(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.String(http.StatusOK, "Successfully Added proxy server")
}
