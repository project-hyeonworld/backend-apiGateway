package common

import (
	model "way-manager/api/shared/common/model"
)

type IController interface {
	Add(proxyServer *model.ProxyServer) error
}
