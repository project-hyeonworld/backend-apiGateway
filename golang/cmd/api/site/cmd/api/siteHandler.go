package site

import (
	"context"
	"net/http"
	"time"
	model "way-manager/api/shared/common/model"
	pb "way-manager/api/shared/proto/nginx"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	nginxHdlr pb.NginxHandlerClient
	ctrl      IController
}

func NewHandler(nginxHdlr pb.NginxHandlerClient, ctrl IController) *Handler {
	return &Handler{nginxHdlr: nginxHdlr, ctrl: ctrl}
}

func (h *Handler) Add(c *gin.Context) {
	var input model.ProxyServer
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.ctrl.Add(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.restartNginx(); err != nil {
		c.String(http.StatusInternalServerError, "ngix not restarted :%v", err)
		return
	}

	c.String(http.StatusOK, "Successfully Added proxy server")
}

func (h *Handler) restartNginx() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)

	_, err := h.nginxHdlr.Restart(ctx, &pb.RestartRequest{})
	if err != nil {
		return err
	}
	defer cancel()
	return nil
}
