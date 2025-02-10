package nginx

import (
	"context"
	"net/http"
	pb "way-manager/api/shared/proto/nginx"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	pb.UnimplementedNginxHandlerServer
	ctrl IController
}

func NewHandler(ctrl IController) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) Restart(ctx context.Context, restartRequest *pb.RestartRequest) (*pb.RestartResponse, error) {
	h.ctrl.Restart()
	return &pb.RestartResponse{Success: true}, nil
}

func (h *Handler) RestartApi(c *gin.Context) {
	h.ctrl.Restart()
	c.String(http.StatusOK, "Successfully Added proxy server")
}
