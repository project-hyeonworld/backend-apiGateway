package site

import (
	"context"
	"net/http"
	"time"
	"way-manager/api/shared/common/model"
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

// Add
// @Summary Add a new proxy server.
// @Description Adds a new proxy server by accepting parameters via query and restarting nginx and restarting nginx by calling restartNginx().
// @Tags site
// @Accept  json
// @Produce  json
// @Param proxyServer query model.ProxyServer true "Proxy server details"
// @Success 200 {string} string "Successfully Added proxy server"
// @Failure 400 {object} map[string]string "Invalid request parameters"}
// @Failure 500 {object} map[string]string "Internal Server Error"}
// @Router /site [post]
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

// restartNginx
// @Summary Restart Nginx system.
// @Description Attempts to restart the Nginx service by calling the Restart method on the Nginx handler with a context timeout of 3 minutes.
// @Tags site
// @Success 200 {string} string "Nginx service restarted successfully"
// @Failure 500 {object} map[string]string "Failed to restart Nginx"}
// @Router /api/nginx/restart [post]
func (h *Handler) restartNginx() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)

	_, err := h.nginxHdlr.Restart(ctx, &pb.RestartRequest{})
	if err != nil {
		return err
	}
	defer cancel()
	return nil
}
