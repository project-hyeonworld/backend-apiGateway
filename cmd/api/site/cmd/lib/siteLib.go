package lib

import (
	"fmt"
	di "way-manager/api/site/cmd/configuration/dependencyInjection"
	"way-manager/api/site/cmd/configuration/secret"
	commonSecret "way-manager/configuration/secret"

	"github.com/gin-gonic/gin"
)

func StartApiServer(commonSecretValue *commonSecret.Value, router *gin.Engine) error {
	secretValue := secret.Value{}
	if err := secretValue.Init(commonSecretValue); err != nil {
		return fmt.Errorf("failed to initialize secret value: %w", err)
	}

	container := di.Container{}
	if err := container.Init(&secretValue); err != nil {
		return fmt.Errorf("failed to initialize container: %w", err)
	}

	siteGroup := router.Group("/site")
	{
		siteGroup.POST("", container.SiteHandler.Add)
	}
	return nil
}
