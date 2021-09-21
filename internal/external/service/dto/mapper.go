package dto

import (
	"errors"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	"strconv"
	"strings"
)

func MapServiceConfigToExternal(config model.ServiceConfig) (*Service, error) {
	var (
		ports        []string
		volumes      []string
		healthChecks []*HealthCheck
	)
	for _, port := range config.PortBindings() {
		ports = append(ports, port.String())
	}
	for _, volume := range config.VolumeBindings() {
		volumes = append(volumes, volume.String())
	}
	for _, healthCheck := range config.HealthChecks() {
		healthChecks = append(healthChecks, &HealthCheck{
			Type:     string(healthCheck.Type()),
			Port:     healthCheck.Port(),
			Endpoint: healthCheck.Endpoint(),
		})
	}

	return &Service{
		FilePath:     config.FilePath(),
		Image:        config.Image(),
		DomainName:   config.DomainName(),
		Environment:  config.Environments(),
		Ports:        ports,
		Volumes:      volumes,
		HealthChecks: healthChecks,
		Networks:     config.Networks(),
		Tag:          string(config.Tag()),
	}, nil
}

func MapExternalToServiceConfig(name string, svc *Service) (model.ServiceConfig, error) {
	var (
		ports          []model.Port
		volumeBindings []model.Volume
		healthChecks   []model.HealthCheck
	)

	for _, port := range svc.Ports {
		portProto := strings.SplitN(port, "/", 2)
		if len(portProto) < 1 || len(portProto) > 2 {
			return nil, errors.New("error [MapExternalToServiceConfig]: invalid port and protocol")
		}
		proto := ""
		if len(portProto) == 2 {
			proto = portProto[1]
		}
		p := strings.SplitN(portProto[0], ":", 2)
		var pInt []int
		for _, v := range p {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			pInt = append(pInt, vInt)
		}
		switch len(pInt) {
		case 2:
			switch proto {
			case string(model.ProtocolTCP):
				ports = append(ports, model.NewPortBindingTCP(pInt[0], pInt[1]))
			case string(model.ProtocolUDP):
				ports = append(ports, model.NewPortBindingUDP(pInt[0], pInt[1]))
			default:
				ports = append(ports, model.NewPortBinding(pInt[0], pInt[1]))
			}
		case 1:
			ports = append(ports, model.NewPort(pInt[0]))
		}
	}
	for _, volume := range svc.Volumes {
		v := strings.SplitN(volume, ":", 2)
		switch len(v) {
		case 2:
			volumeBindings = append(volumeBindings, model.NewVolumeBinding(v[0], v[1]))
		case 1:
			volumeBindings = append(volumeBindings, model.NewVolume(v[0]))
		}
	}
	for _, healthCheck := range svc.HealthChecks {
		switch healthCheck.Type {
		case string(model.HealthCheckHTTP):
			hc := model.NewHealthCheckHTTP(healthCheck.Port, healthCheck.Endpoint)
			healthChecks = append(healthChecks, hc)
		case string(model.HealthCheckTCP):
			hc := model.NewHealthCheckTCP(healthCheck.Port)
			healthChecks = append(healthChecks, hc)
		}
	}

	if svc.FilePath != "" {
		config := model.NewCustomServiceConfig(name, svc.DomainName, svc.FilePath, ports)
		return config, nil
	}

	return model.NewServiceConfig(
		name,
		svc.DomainName,
		svc.Image,
		svc.Environment,
		ports,
		volumeBindings,
		healthChecks,
		svc.Networks,
		model.ServiceTag(svc.Tag),
	), nil
}
