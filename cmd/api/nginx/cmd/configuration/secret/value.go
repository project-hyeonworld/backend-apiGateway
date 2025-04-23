package configuration_secret

import (
	commonSecret "way-manager/configuration/secret"
)

type Value struct {
	CommonValue commonSecret.Value
}

func (v *Value) Init(commonValue *commonSecret.Value) error {
	v.CommonValue = getCommonScretValue(commonValue)
	return nil
}

func getCommonScretValue(commonValue *commonSecret.Value) commonSecret.Value {
	if commonValue != nil {
		return *commonValue
	}
	CommonValue := commonSecret.Value{}
	CommonValue.NginxApiIp = commonSecret.LOCALHOST
	CommonValue.NginxApiPort = 5001
	return CommonValue
}
