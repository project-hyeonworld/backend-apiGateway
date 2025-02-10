package site

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	model "way-manager/api/shared/common/model"
	"way-manager/api/site/cmd/configuration/secret"
)

type NginxConfigBusiness struct {
	application  map[string]secret.ApplicationInfo
	availableDir string
	enabledDir   string
}

func NewNginxConfigBusiness(secretValue *secret.Value) NginxConfigBusiness {
	return NginxConfigBusiness{
		application:  secretValue.Applications,
		availableDir: secretValue.SiteValue.AvailableDir,
		enabledDir:   secretValue.SiteValue.EnabledDir,
	}
}

func (biz NginxConfigBusiness) CreateSymlink(applicationName *string) error {

	if applicationName != nil {
		expectedFile := filepath.Join(biz.getSymlinkFileName(applicationName))
		if _, err := os.Stat(expectedFile); err == nil {
			// File exists, so we don't need to create a new symlink
			fmt.Printf("Configuration for %s already exists", *applicationName)
			return nil
		}
	}

	confFiles, err := filepath.Glob(biz.getFileName(applicationName))
	if err != nil {
		return fmt.Errorf("failed to get conf files: %v", err)
	}
	for _, confFile := range confFiles {
		fileName := filepath.Base(confFile)
		symlinkPath := filepath.Join(biz.getSymlinkFileName(applicationName))

		err := os.Symlink(confFile, symlinkPath)
		if err != nil {
			return fmt.Errorf("failed to create symlink for %s: %v", fileName, err)
		}
		fmt.Printf("Created symlink for %s\n", fileName)
	}
	return nil
}
func (biz *NginxConfigBusiness) PatchFile(applicationName, content *string) error {
	filename := biz.getFileName(applicationName)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(*content)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func (biz *NginxConfigBusiness) AddProxyServer(config *model.NginxConfig, proxyServer *model.ProxyServer) {
	config.AddUpstream(proxyServer)
}

func (biz *NginxConfigBusiness) getFileName(applicationName *string) string {
	return biz.availableDir + "/" + *applicationName + ".conf"
}

func (biz *NginxConfigBusiness) getSymlinkFileName(applicationName *string) string {
	return biz.enabledDir + "/" + *applicationName + ".conf"
}

func (biz *NginxConfigBusiness) CreateFile(proxyServer *model.ProxyServer) (string, error) {
	filename := biz.getFileName(&proxyServer.ApplicationName)
	configContent, err := biz.createConfigString(proxyServer)
	if err != nil {
		return "", fmt.Errorf("failed to get config content: %v", err)
	}
	file, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(configContent)
	if err != nil {
		return "", fmt.Errorf("error writing to file: %v", err)
	}

	return filename, nil
}

func (biz *NginxConfigBusiness) createConfigString(proxyServer *model.ProxyServer) (string, error) {
	serverLocationPath := biz.getServerLocationPath(&proxyServer.ApplicationName)
	config := model.NginxConfig{}
	config.Fill(proxyServer, &serverLocationPath)
	return config.ToString(), nil
}

func (biz *NginxConfigBusiness) getServerLocationPath(applicationName *string) string {
	return biz.application[*applicationName].ApiLocation
}

func (biz *NginxConfigBusiness) parseApplicationName(applicationName *string) (string, string) {
	parts := strings.SplitN(*applicationName, "_", 2)
	projectName := parts[0]

	lastPart := parts[1]
	lastHyphenIndex := strings.LastIndex(lastPart, "-")
	serviceName := lastPart[:lastHyphenIndex]
	return projectName, serviceName
}

func (biz *NginxConfigBusiness) ReadFile(applicationName *string) (string, error) {
	filename := biz.getFileName(applicationName)
	content, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("file does not exist: %w", err)
		}
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(content), nil
}

func (biz *NginxConfigBusiness) ParseNginxConfig(content *string) (model.NginxConfig, error) {
	config := model.NginxConfig{}

	lines := strings.Split(*content, "\n")

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "upstream") {
			upstreamBlock, newIndex := parseUpstreamBlcok(lines, i)
			config.Upstream = upstreamBlock
			i = newIndex
		}
		if strings.HasPrefix(line, "server") {
			serverBlock, newIndex := parseServerBlock(lines, i)
			config.Server = serverBlock
			i = newIndex
		}
	}
	return config, nil
}

func parseUpstreamBlcok(lines []string, startIndex int) (model.UpstreamBlock, int) {
	block := model.UpstreamBlock{}
	for i := startIndex; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "}" {
			return block, i
		}
		if strings.HasPrefix(line, "upstream") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				block.Name = parts[1]
			}
			continue
		}
		if line == "ip_hash;" {
			block.Method = "ip_hash"
			continue
		}
		if strings.HasPrefix(line, "server") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				serverInfo := strings.TrimSuffix(parts[1], ";")
				ipPort := strings.Split(serverInfo, ":")
				if len(ipPort) == 2 {
					block.Ip = append(block.Ip, ipPort[0])
					if port, err := strconv.ParseUint(ipPort[1], 10, 16); err == nil {
						block.Port = append(block.Port, uint16(port))
					}
				}
			}
		}
	}
	return block, len(lines) - 1
}

func parseServerBlock(lines []string, startIndex int) (model.ServerBlock, int) {
	block := model.ServerBlock{
		Locations:    []model.LocationBlock{},
		ProxyHeaders: make(map[string]string),
	}
	inLocation := false
	var currentLocation model.LocationBlock

	for i := startIndex; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "}" {
			if inLocation {
				block.Locations = append(block.Locations, currentLocation)
				inLocation = false
			} else {
				return block, i
			}
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		if parts[0] == "listen" {
			if port, err := strconv.Atoi(strings.TrimSuffix(parts[1], ";")); err == nil {
				block.ListenPort = port
			}
			continue
		}

		if parts[0] == "location" {
			if inLocation {
				block.Locations = append(block.Locations, currentLocation)
			}
			inLocation = true
			currentLocation = model.LocationBlock{Path: parts[1]}
		}

		if parts[0] == "proxy_pass" {
			if inLocation {
				currentLocation.ProxyPass = strings.TrimSuffix(parts[1], ";")
			}
		}
		if parts[0] == "proxy_set_header" {
			if len(parts) >= 3 {
				key := parts[1]
				value := strings.TrimSuffix(strings.Join(parts[2:], " "), ";")
				block.ProxyHeaders[key] = value
			}
		}
	}

	if inLocation {
		block.Locations = append(block.Locations, currentLocation)
	}

	return block, len(lines) - 1
}
