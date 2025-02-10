package site

import (
	model "way-manager/api/shared/common/model"
)

type Controller struct {
	svc IService
}

func NewController(svc IService) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) Add(proxyServer *model.ProxyServer) error {
	if err := c.svc.Add(proxyServer); err != nil {
		return err
	}
	return nil
}
