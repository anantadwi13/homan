package domain

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Config interface {
	// BasePath contains full path to base directory
	BasePath() string
	// ConfigPath contains full path
	ConfigPath() string
	// CustomPath contains full path
	CustomPath() string
	// DataPath contains full path
	DataPath() string
	// TempPath contains full path
	TempPath() string

	PublicIP() string
	ProjectName() string
	ServiceRegistryConfPath() string
	SystemNamePrefix() string

	DaemonPort() int
}

type ConfigParams struct {
	BasePath                string // BasePath directory of project, default "$(pwd)/.homan"
	ConfigPath              string // ConfigPath relative to BasePath, default "/config"
	DataPath                string // DataPath relative to BasePath, default "/data"
	CustomPath              string // CustomPath relative to BasePath, default "/custom"
	TempPath                string // TempPath relative to BasePath, default "/temp"
	PublicIP                string // PublicIP default : Public IP of current server, example "45.45.45.45"
	ProjectName             string // ProjectName default "homan"
	ServiceRegistryConfName string // ServiceRegistryConfName relative to ConfigPath, default "registry"
	SystemNamePrefix        string // default "system-"
	DaemonPort              int    // default 5555
}

type config struct {
	basePath                string
	configPath              string
	customPath              string
	dataPath                string
	tempPath                string
	publicIP                string
	projectName             string
	serviceRegistryConfPath string
	systemNamePrefix        string
	daemonPort              int
}

func NewConfig(params ConfigParams) (Config, error) {
	c := &config{}

	if params.BasePath != "" {
		c.basePath = filepath.Clean(params.BasePath)
	} else {
		dir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		c.basePath = filepath.Join(dir, "/.homan")
	}

	joinPath := func(first, second, defaultSecond string) string {
		if second != "" {
			return filepath.Join(first, second)
		} else {
			return filepath.Join(first, defaultSecond)
		}
	}
	checkEmpty := func(value, defaultValue string) string {
		if value != "" {
			return value
		} else {
			return defaultValue
		}
	}

	c.configPath = joinPath(c.basePath, params.ConfigPath, "/config")
	c.customPath = joinPath(c.basePath, params.CustomPath, "/custom")
	c.dataPath = joinPath(c.basePath, params.DataPath, "/data")
	c.tempPath = joinPath(c.basePath, params.TempPath, "/temp")

	publicIp := "127.0.0.1"
	publicIpRes, err := http.Get("https://api.ipify.org?format=text")
	if err == nil {
		ip, err := io.ReadAll(publicIpRes.Body)
		if err == nil {
			publicIp = string(ip)
		}
	}

	c.publicIP = checkEmpty(params.PublicIP, publicIp)
	c.projectName = checkEmpty(params.ProjectName, "homan")
	c.serviceRegistryConfPath = joinPath(c.configPath, params.ServiceRegistryConfName, "registry")
	c.systemNamePrefix = checkEmpty(params.SystemNamePrefix, "system-")

	if params.DaemonPort < 0 || params.DaemonPort >= 65536 {
		return nil, errors.New("DaemonPort is invalid")
	} else if params.DaemonPort == 0 {
		c.daemonPort = 5555
	} else {
		c.daemonPort = params.DaemonPort
	}

	return c, nil
}

func (c *config) BasePath() string {
	return c.basePath
}

func (c *config) ConfigPath() string {
	return c.configPath
}

func (c *config) CustomPath() string {
	return c.customPath
}

func (c *config) DataPath() string {
	return c.dataPath
}

func (c *config) TempPath() string {
	return c.tempPath
}

func (c *config) PublicIP() string {
	return c.publicIP
}

func (c *config) ProjectName() string {
	return c.projectName
}

func (c *config) ServiceRegistryConfPath() string {
	return c.serviceRegistryConfPath
}

func (c *config) SystemNamePrefix() string {
	return c.systemNamePrefix
}

func (c *config) DaemonPort() int {
	return c.daemonPort
}
