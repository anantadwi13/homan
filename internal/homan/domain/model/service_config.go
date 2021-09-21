package model

import "path/filepath"

type ServiceTag string

const (
	TagWeb     = ServiceTag("website")
	TagGateway = ServiceTag("gateway")
	TagProxy   = ServiceTag("proxy")
	TagDNS     = ServiceTag("dns")
	TagCertMan = ServiceTag("certman")
	TagDB      = ServiceTag("database")
)

type ServiceConfig interface {
	Name() string       // required
	DomainName() string // required
	FilePath() string
	Image() string
	Environments() []string
	PortBindings() []Port
	VolumeBindings() []Volume
	HealthChecks() []HealthCheck
	Networks() []string
	Tag() ServiceTag

	IsCustom() bool
	Validator
}

type sc struct {
	file         string
	domainName   string
	image        string
	name         string
	environments []string
	ports        []Port
	volBindings  []Volume
	healthChecks []HealthCheck
	networks     []string
	serviceTag   ServiceTag
	isCustom     bool
}

func NewServiceConfig(
	name string, domainName string, image string, environments []string, portBindings []Port, volBindings []Volume,
	healthChecks []HealthCheck, networks []string, serviceTag ServiceTag,
) ServiceConfig {
	return &sc{
		image:        image,
		name:         name,
		domainName:   domainName,
		environments: environments,
		ports:        portBindings,
		volBindings:  volBindings,
		healthChecks: healthChecks,
		networks:     networks,
		isCustom:     false,
		serviceTag:   serviceTag,
	}
}

func NewCustomServiceConfig(name string, domainName string, filePath string, portBindings []Port) ServiceConfig {
	return &sc{
		name:       name,
		file:       filepath.Clean(filePath),
		domainName: domainName,
		ports:      portBindings,
		isCustom:   true,
		serviceTag: TagWeb,
	}
}

func (s *sc) FilePath() string {
	return s.file
}

func (s *sc) Image() string {
	return s.image
}

func (s *sc) DomainName() string {
	return s.domainName
}

func (s *sc) Name() string {
	return s.name
}

func (s *sc) Environments() []string {
	return s.environments
}

func (s *sc) PortBindings() []Port {
	return s.ports
}

func (s *sc) VolumeBindings() []Volume {
	return s.volBindings
}

func (s *sc) HealthChecks() []HealthCheck {
	return s.healthChecks
}

func (s *sc) Networks() []string {
	return s.networks
}

func (s *sc) Tag() ServiceTag {
	return s.serviceTag
}

func (s *sc) IsCustom() bool {
	return s.isCustom
}

func (s *sc) IsValid() bool {
	if s.isCustom {
		return s.file != ""
	} else {
		for _, port := range s.ports {
			if port == nil {
				return false
			}
			if !port.IsValid() {
				return false
			}
		}
		for _, volBinding := range s.volBindings {
			if volBinding == nil {
				return false
			}
			if !volBinding.IsValid() {
				return false
			}
		}
		for _, hc := range s.healthChecks {
			if hc == nil {
				return false
			}
			if !hc.IsValid() {
				return false
			}
		}

		return s.image != "" && s.name != ""
	}
}
