package secret

import (
	commonSecret "way-manager/configuration/secret"
)

type ProjectInfo struct {
	Port uint16
}
type ApplicationInfo struct {
	ApiLocation string
}

type SiteValue struct {
	AvailableDir string
	EnabledDir   string
}

type Value struct {
	CommonValue  commonSecret.Value
	Projects     map[string]ProjectInfo
	Applications map[string]ApplicationInfo
	SiteValue    SiteValue
}

func (v *Value) Init(commonSecretValue *commonSecret.Value) error {
	v.CommonValue = getCommonScretValue(commonSecretValue)
	v.Projects = map[string]ProjectInfo{
		"{GO_ENVIRONMENT_NAME}": {
			Port: {MY_BACKEND_PORT},
		},
	}
	v.Applications = map[string]ApplicationInfo{
		"{MY_SESSION_APPLICATION_NAME}": {
			ApiLocation: "{MY_SESSION_APPLICATION_API_LOCATION}",
		},
		"{MY_USER_APPLICATION_NAME}": {
			ApiLocation: "{MY_USER_APPLICATION_API_LOCATION}",
		},
	}
	v.SiteValue = SiteValue{
		AvailableDir: "{AVAILABLE_DIRECTORY}",
		EnabledDir:   "{ENABLED_DIRECTORY}",
	}
	return nil
}

func getCommonScretValue(commonSecretValue *commonSecret.Value) commonSecret.Value {
	if commonSecretValue != nil {
		return *commonSecretValue
	}
	CommonValue := commonSecret.Value{}
	CommonValue.NginxApiIp = commonSecret.LOCALHOST
	CommonValue.NginxApiPort = 5001
	return CommonValue
}
