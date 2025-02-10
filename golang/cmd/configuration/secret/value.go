package configuration_secret

type Value struct {
	SiteApiIp    string
	SiteApiPort  uint16
	NginxApiIp   string
	NginxApiPort uint16
}

const (
	LOCALHOST = "localhost"
)

func (v *Value) Init() {
	v.NginxApiIp = LOCALHOST
	v.NginxApiPort = 5001
}
