package common

import (
	model "way-manager/api/shared/common/model"
)

type IService interface {
	Add(proxyServer *model.ProxyServer) error
}
