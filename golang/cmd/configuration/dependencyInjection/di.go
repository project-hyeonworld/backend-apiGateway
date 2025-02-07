package dependencyinjection

import (
	common "way-manager/api/shared/common"
	secret "way-manager/configuration/secret"

	site "way-manager/api/site/cmd/api"
)

type Container struct {
	siteRepository common.IRepository
	siteService    common.IService
	siteController common.IController
	siteHandler    common.IHandler
}

func (c *Container) Init(secretValue secret.Value) {
	c.siteRepository = site.NewRepository()
	c.siteService = site.NewService(c.siteRepository)
	c.siteController = site.NewController(c.siteService)
}
