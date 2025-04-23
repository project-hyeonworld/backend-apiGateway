package site

import (
	"way-manager/api/shared/common"
	model "way-manager/api/shared/common/model"
)

type IService interface {
	common.IService
	Add(proxyServer *model.ProxyServer) error
}
