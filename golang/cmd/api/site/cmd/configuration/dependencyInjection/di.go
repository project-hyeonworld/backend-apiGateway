package dependencyinjection

import (
	"way-manager/api/shared/common"
	site "way-manager/api/site/cmd/api"
	"way-manager/api/site/cmd/configuration/secret"
)

type Container struct {
	NginxConfigBusiness site.NginxConfigBusiness

	SiteRepository common.IRepository
	SiteService    common.IService
	SiteController common.IController
	SiteHandler    common.IHandler
}

func (c *Container) Init(secretValue *secret.Value) {
	c.NginxConfigBusiness = site.NewNginxConfigBusiness(secretValue)

	c.SiteRepository = site.NewRepository()
	c.SiteService = site.NewService(c.SiteRepository, c.NginxConfigBusiness)
	c.SiteController = site.NewController(c.SiteService)

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", secretValue.CommonValue.NginxApiIp, secretValue.CommonValue.NginxApiPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Errorf("failed to connect to nginx client %v", err)
	}
	nginxHandler := pb.NewNginxHandlerClient(conn)
	c.SiteHandler = site.NewHandler(nginxHandler, c.SiteController)
}
