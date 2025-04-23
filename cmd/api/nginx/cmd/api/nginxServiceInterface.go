package nginx

import (
	"way-manager/api/shared/common"
)

type IService interface {
	common.IService
	Restart() error
}
