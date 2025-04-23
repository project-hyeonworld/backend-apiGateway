package dependencyinjection

import (
	"fmt"
	"net"
	nginx "way-manager/api/nginx/cmd/api"
	secret "way-manager/api/nginx/cmd/configuration/secret"
	nginxPb "way-manager/api/shared/proto/nginx"

	"google.golang.org/grpc"
)

type Container struct {
	OsBusiness nginx.OsBusiness

	NginxService    nginx.IService
	NginxController nginx.IController
	NginxHandler    nginx.IHandler
}

func (c *Container) Init(secretValue *secret.Value) {
	c.OsBusiness = nginx.NewOsBusiness()
	c.NginxService = nginx.NewService(c.OsBusiness)
	c.NginxController = nginx.NewController(c.NginxService)
	c.NginxHandler = nginx.NewHandler(c.NginxController)

	go startGrpcServer(secretValue, c.NginxHandler)
	a := 4
	fmt.Printf("%d", a)
}

func startGrpcServer(secretValue *secret.Value, handler nginx.IHandler) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", secretValue.CommonValue.NginxApiIp, secretValue.CommonValue.NginxApiPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	s := grpc.NewServer()

	nginxPb.RegisterNginxHandlerServer(s, handler)
	if servErr := s.Serve(lis); servErr != nil {
		return fmt.Errorf("failed to server: %w", err)
	}
	return nil
}
