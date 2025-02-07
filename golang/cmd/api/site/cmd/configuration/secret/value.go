package secret

type ApplicationInfo struct {
	ApiLocation string
}

type Value struct {
	Applications map[string]ApplicationInfo
}

func (v *Value) Init() error {
	v.Applications = map[string]ApplicationInfo{
		"{MY_SESSION_APPLICATION_NAME}": {
			ApiLocation: "{MY_SESSION_APPLICATION_API_LOCATION}",
		},
	}
	return nil
}
