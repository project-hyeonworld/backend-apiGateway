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
	c.SiteHandler = site.NewHandler(c.SiteController)
}
