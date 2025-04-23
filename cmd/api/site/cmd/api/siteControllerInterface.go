package site

import (
	"way-manager/api/shared/common"
	model "way-manager/api/shared/common/model"
)

type IController interface {
	common.IController
	Add(proxyServer *model.ProxyServer) error
}
