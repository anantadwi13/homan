package dto

import (
	"errors"
	model2 "github.com/anantadwi13/homan/internal/homan/domain/model"
	"strconv"
	"strings"
)

func MapServiceConfigToExternal(config model2.ServiceConfig) (*Service, error) {
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

func MapExternalToServiceConfig(name string, svc *Service) (model2.ServiceConfig, error) {
	var (
		ports          []model2.Port
		volumeBindings []model2.Volume
		healthChecks   []model2.HealthCheck
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
			case string(model2.ProtocolTCP):
				ports = append(ports, model2.NewPortBindingTCP(pInt[0], pInt[1]))
			case string(model2.ProtocolUDP):
				ports = append(ports, model2.NewPortBindingUDP(pInt[0], pInt[1]))
			default:
				ports = append(ports, model2.NewPortBinding(pInt[0], pInt[1]))
			}
		case 1:
			ports = append(ports, model2.NewPort(pInt[0]))
		}
	}
	for _, volume := range svc.Volumes {
		v := strings.SplitN(volume, ":", 2)
		switch len(v) {
		case 2:
			volumeBindings = append(volumeBindings, model2.NewVolumeBinding(v[0], v[1]))
		case 1:
			volumeBindings = append(volumeBindings, model2.NewVolume(v[0]))
		}
	}
	for _, healthCheck := range svc.HealthChecks {
		switch healthCheck.Type {
		case string(model2.HealthCheckHTTP):
			hc := model2.NewHealthCheckHTTP(healthCheck.Port, healthCheck.Endpoint)
			healthChecks = append(healthChecks, hc)
		case string(model2.HealthCheckTCP):
			hc := model2.NewHealthCheckTCP(healthCheck.Port)
			healthChecks = append(healthChecks, hc)
		}
	}

	if svc.FilePath != "" {
		config := model2.NewCustomServiceConfig(name, svc.DomainName, svc.FilePath, ports)
		return config, nil
	}

	return model2.NewServiceConfig(
		name,
		svc.DomainName,
		svc.Image,
		svc.Environment,
		ports,
		volumeBindings,
		healthChecks,
		svc.Networks,
		model2.ServiceTag(svc.Tag),
	), nil
}
