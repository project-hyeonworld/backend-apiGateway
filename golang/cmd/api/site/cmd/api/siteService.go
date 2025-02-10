package site

import (
	"fmt"
	model "way-manager/api/shared/common/model"
)

type Service struct {
	repo IRepository
	biz  NginxConfigBusiness
}

func NewService(repo IRepository, biz NginxConfigBusiness) *Service {
	return &Service{repo: repo, biz: biz}
}

func (s *Service) Add(proxyServer *model.ProxyServer) error {
	content, _ := s.biz.ReadFile(&proxyServer.ApplicationName)
	if content == "" {
		s.biz.CreateFile(proxyServer)
		s.biz.CreateSymlink(&proxyServer.ApplicationName)
	}
	config, err := s.biz.ParseNginxConfig(&content)
	if err != nil {
		return err
	}
	s.biz.AddProxyServer(&config, proxyServer)

	patchContent := config.ToString()

	if err := s.biz.PatchFile(&proxyServer.ApplicationName, &patchContent); err != nil {
		return fmt.Errorf("failed to patch file: %v\n", err)
	}

	return nil
}
