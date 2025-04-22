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

// Restart
// @Deprecated This endpoint is not used, more research required.
// @Summary Restart the Nginx using gRPC.
// @Description Restarts the API by calling the Restart method in the controller.
// @Tags Nginx
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Successfully Added proxy server"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/restart [patch]
func (h *Handler) Restart(ctx context.Context, restartRequest *pb.RestartRequest) (*pb.RestartResponse, error) {
	h.ctrl.Restart()
	return &pb.RestartResponse{Success: true}, nil
}

// RestartApi
// @Summary Restart the Nginx using plain API.
// @Description Restarts the Nginx by calling the Restart method in the controller.
// @Tags Nginx
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Successfully Added proxy server"
// @Failure 500 {string} string "Internal Server Error"
// @Router /nginx/restart [patch]
func (h *Handler) RestartApi(c *gin.Context) {
	h.ctrl.Restart()
	c.String(http.StatusOK, "Successfully Added proxy server")
}
