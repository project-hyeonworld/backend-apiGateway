package nginx

import (
	"way-manager/api/shared/common"
)

type IController interface {
	common.IController
	Restart() error
}
