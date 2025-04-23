package model

type ProxyServer struct {
	ApplicationName string `form:"applicationName"`
	Ip              string `form:"ip"`
	Port            uint16 `form:"port"`
}
