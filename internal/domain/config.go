package domain

import (
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

	ProjectName() string
	ServiceRegistryConfPath() string
	SystemNamePrefix() string
}

type ConfigParams struct {
	BasePath                string // BasePath directory of project, default "$(pwd)/.homan"
	ConfigPath              string // ConfigPath relative to BasePath, default "/config"
	DataPath                string // DataPath relative to BasePath, default "/data"
	CustomPath              string // CustomPath relative to BasePath, default "/custom"
	TempPath                string // TempPath relative to BasePath, default "/temp"
	ProjectName             string // ServiceRegistryConfName relative to ConfigPath, default "homan"
	ServiceRegistryConfName string // ServiceRegistryConfName relative to ConfigPath, default "registry"
	SystemNamePrefix        string // default "system-"
}

type config struct {
	basePath                string
	configPath              string
	customPath              string
	dataPath                string
	tempPath                string
	projectName             string
	serviceRegistryConfPath string
	systemNamePrefix        string
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
	c.projectName = checkEmpty(params.ProjectName, "homan")
	c.serviceRegistryConfPath = joinPath(c.configPath, params.ServiceRegistryConfName, "registry")
	c.systemNamePrefix = checkEmpty(params.SystemNamePrefix, "system-")

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

func (c *config) ProjectName() string {
	return c.projectName
}

func (c *config) ServiceRegistryConfPath() string {
	return c.serviceRegistryConfPath
}

func (c *config) SystemNamePrefix() string {
	return c.systemNamePrefix
}
