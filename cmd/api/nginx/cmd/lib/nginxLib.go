package lib

import (
	"fmt"
	di "way-manager/api/nginx/cmd/configuration/dependencyInjection"
	secret "way-manager/api/nginx/cmd/configuration/secret"
	commonSecret "way-manager/configuration/secret"

	"github.com/gin-gonic/gin"
)

func StartApiServer(commonSecretValue *commonSecret.Value, router *gin.Engine) error {
	secretValue := secret.Value{}
	if err := secretValue.Init(commonSecretValue); err != nil {
		return fmt.Errorf("failed to initialize secret value: %w", err)
	}
	container := di.Container{}
	container.Init(&secretValue)

	nginxGroup := router.Group("/nginx")
	{
		nginxGroup.PATCH("/restart", container.NginxHandler.RestartApi)
	}
	return nil
}

func StartGrpcServer(commonSecretValue *commonSecret.Value) error {
	return nil
}
