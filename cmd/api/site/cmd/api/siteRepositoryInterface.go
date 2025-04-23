package site

import (
	"way-manager/api/shared/common"
	model "way-manager/api/shared/common/model"
)

type IRepository interface {
	common.IRepository
	Add(proxyServer *model.ProxyServer) error
}
