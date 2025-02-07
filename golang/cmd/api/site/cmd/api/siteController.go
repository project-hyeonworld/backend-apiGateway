package site

import (
	common "way-manager/api/shared/common"
	model "way-manager/api/shared/common/model"
)

type Controller struct {
	svc common.IService
}

func NewController(svc common.IService) *Controller {
	return &Controller{svc: svc}
}

func (c *Controller) Add(proxyServer *model.ProxyServer) error {
	if err := c.svc.Add(proxyServer); err != nil {
		return err
	}
	return nil
}
