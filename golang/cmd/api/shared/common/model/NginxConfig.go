package model

import (
	"fmt"
	"strings"
)

type UpstreamBlock struct {
	Name   string
	Method string
	Ip     []string
	Port   []uint16
}

type ServerBlock struct {
	ListenPort   uint16
	Locations    []LocationBlock
	ProxyHeaders map[string]string
}

type LocationBlock struct {
	Path      string
	ProxyPass string
}

type NginxConfig struct {
	Upstream UpstreamBlock
	Server   ServerBlock
}

func (c *NginxConfig) Fill(proxyServer *ProxyServer, serverLocationPath *string, listenPort uint16) *NginxConfig {
	// Modify Upstream
	c.Upstream.Name = proxyServer.ApplicationName
	c.Upstream.Method = "ip_hash"
	c.Upstream.Ip = []string{proxyServer.Ip}
	c.Upstream.Port = []uint16{proxyServer.Port}

	// Modify Server
	c.Server.ListenPort = listenPort
	c.Server.Locations = []LocationBlock{
		{
			Path:      *serverLocationPath,
			ProxyPass: "http://" + proxyServer.ApplicationName,
		},
	}
	c.Server.ProxyHeaders = map[string]string{
		"Host":              "$host",
		"X-Real-IP":         "$remote_addr",
		"X-Forwarded-For":   "$proxy_add_x_forwarded_for",
		"X-Forwarded-Proto": "$scheme",
	}

	return c
}

func (c *NginxConfig) AddUpstream(proxyServer *ProxyServer) {
	c.Upstream.Ip = append(c.Upstream.Ip, proxyServer.Ip)
	c.Upstream.Port = append(c.Upstream.Port, proxyServer.Port)
}

func (c *NginxConfig) ToString() string {
	var result strings.Builder

	// Write upstream block
	result.WriteString(fmt.Sprintf("upstream %s {\n", c.Upstream.Name))
	result.WriteString(fmt.Sprintf("\t%s;\n", c.Upstream.Method))
	for i, ip := range c.Upstream.Ip {
		port := c.Upstream.Port[i]
		result.WriteString(fmt.Sprintf("\tserver %s:%d;\n", ip, port))
	}
	result.WriteString("}\n\n")

	// Write server block
	result.WriteString("server {\n")
	result.WriteString(fmt.Sprintf("\tlisten %d;\n", c.Server.ListenPort))

	for _, location := range c.Server.Locations {
		result.WriteString(fmt.Sprintf("\tlocation %s {\n", location.Path))
		result.WriteString(fmt.Sprintf("\t\tproxy_pass %s;\n", location.ProxyPass))
		result.WriteString("\t}\n")
	}

	for key, value := range c.Server.ProxyHeaders {
		result.WriteString(fmt.Sprintf("\tproxy_set_header %s %s;\n", key, value))
	}

	result.WriteString("}\n")
	return result.String()
}
