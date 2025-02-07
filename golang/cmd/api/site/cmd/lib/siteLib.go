package lib

import (
	"fmt"
	di "way-manager/api/site/cmd/configuration/dependencyInjection"
	"way-manager/api/site/cmd/configuration/secret"

	"github.com/gin-gonic/gin"
)

func Start(router *gin.Engine) error {
	secretValue := secret.Value{}
	if err := secretValue.Init(); err != nil {
		return fmt.Errorf("failed to initialize secret value: %w", err)
	}

	container := di.Container{}
	container.Init(&secretValue)

	siteGroup := router.Group("/site")
	{
		siteGroup.POST("", container.SiteHandler.Add)
	}
	return nil
}
