package dependencyinjection

import (
	"fmt"
	pb "way-manager/api/shared/proto/nginx"
	site "way-manager/api/site/cmd/api"
	"way-manager/api/site/cmd/configuration/secret"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Container struct {
	NginxConfigBusiness site.NginxConfigBusiness

	SiteRepository site.IRepository
	SiteService    site.IService
	SiteController site.IController
	SiteHandler    site.IHandler
}

func (c *Container) Init(secretValue *secret.Value) error {
	c.NginxConfigBusiness = site.NewNginxConfigBusiness(secretValue)
	c.SiteRepository = site.NewRepository()
	c.SiteService = site.NewService(c.SiteRepository, c.NginxConfigBusiness)
	c.SiteController = site.NewController(c.SiteService)

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", secretValue.CommonValue.NginxApiIp, secretValue.CommonValue.NginxApiPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to nginx client %w", err)
	}
	nginxHandler := pb.NewNginxHandlerClient(conn)
	c.SiteHandler = site.NewHandler(nginxHandler, c.SiteController)
}
