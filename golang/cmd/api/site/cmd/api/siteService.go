package site

import (
	common "way-manager/api/shared/common"
	model "way-manager/api/shared/common/model"
)

type Service struct {
	repo common.IRepository
	biz  NginxConfigBusiness
}

func NewService(repo common.IRepository, biz NginxConfigBusiness) *Service {
	return &Service{repo: repo, biz: biz}
}

func (s *Service) Add(proxyServer *model.ProxyServer) error {
	content, err := s.biz.ReadFile(&proxyServer.ApplicationName)
	if content == "" {
		s.biz.CreateFile(proxyServer)
		return nil
	}
	config, err := s.biz.ParseNginxConfig(&content)
	if err != nil {
		return err
	}
	s.biz.AddProxyServer(&config, proxyServer)

	patchContent := config.ToString()
	s.biz.PatchFile(&proxyServer.ApplicationName, &patchContent)
	return nil
}
