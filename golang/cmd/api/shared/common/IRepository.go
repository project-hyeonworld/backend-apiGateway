package common

import (
	model "way-manager/api/shared/common/model"
)

type IRepository interface {
	Add(proxyServer *model.ProxyServer) error
}
