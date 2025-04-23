package nginx

import (
	"context"
	"way-manager/api/shared/common"
	pb "way-manager/api/shared/proto/nginx"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	pb.NginxHandlerServer
	common.IHandler
	Restart(context.Context, *pb.RestartRequest) (*pb.RestartResponse, error)
	RestartApi(c *gin.Context)
}
