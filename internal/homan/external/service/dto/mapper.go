package dto

import (
	domainModel "github.com/anantadwi13/homan/internal/homan/domain/model"
)

func MapServiceConfigToExternal(config domainModel.ServiceConfig) (*Service, error) {
	var (
		ports        []*Port
		volumes      []*Volume
		healthChecks []*HealthCheck
	)
	for _, port := range config.PortBindings() {
		ports = append(ports, &Port{
			HostPort:      port.HostPort(),
			ContainerPort: port.ContainerPort(),
			Protocol:      string(port.Protocol()),
		})
	}
	for _, volume := range config.VolumeBindings() {
		volumes = append(volumes, &Volume{
			HostPath:      volume.HostPath(),
			ContainerPath: volume.ContainerPath(),
			NeedCopy:      volume.NeedCopy(),
		})
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
		IsCustom:     config.IsCustom(),
	}, nil
}

func MapExternalToServiceConfig(name string, svc *Service) (domainModel.ServiceConfig, error) {
	var (
		ports          []domainModel.Port
		volumeBindings []domainModel.Volume
		healthChecks   []domainModel.HealthCheck
	)

	for _, port := range svc.Ports {
		if port.HostPort == domainModel.UnsetPort {
			ports = append(ports, domainModel.NewPort(port.ContainerPort))
			continue
		}
		switch port.Protocol {
		case string(domainModel.ProtocolTCP):
			ports = append(ports, domainModel.NewPortBindingTCP(port.HostPort, port.ContainerPort))
		case string(domainModel.ProtocolUDP):
			ports = append(ports, domainModel.NewPortBindingUDP(port.HostPort, port.ContainerPort))
		default:
			if port.HostPort == domainModel.UnsetPort {
				ports = append(ports, domainModel.NewPort(port.ContainerPort))
			} else {
				ports = append(ports, domainModel.NewPortBinding(port.HostPort, port.ContainerPort))
			}
		}
	}
	for _, volume := range svc.Volumes {
		switch {
		case volume.HostPath == "":
			volumeBindings = append(volumeBindings, domainModel.NewVolume(volume.ContainerPath))
		default:
			if volume.NeedCopy {
				volumeBindings = append(volumeBindings, domainModel.NewVolumeBindingCopy(volume.HostPath, volume.ContainerPath))
			} else {
				volumeBindings = append(volumeBindings, domainModel.NewVolumeBinding(volume.HostPath, volume.ContainerPath))
			}
		}
	}
	for _, healthCheck := range svc.HealthChecks {
		switch healthCheck.Type {
		case string(domainModel.HealthCheckHTTP):
			hc := domainModel.NewHealthCheckHTTP(healthCheck.Port, healthCheck.Endpoint)
			healthChecks = append(healthChecks, hc)
		case string(domainModel.HealthCheckTCP):
			hc := domainModel.NewHealthCheckTCP(healthCheck.Port)
			healthChecks = append(healthChecks, hc)
		}
	}

	if svc.IsCustom {
		config := domainModel.NewCustomServiceConfig(name, svc.DomainName, svc.FilePath, ports)
		return config, nil
	}

	return domainModel.NewServiceConfig(
		name,
		svc.DomainName,
		svc.Image,
		svc.Environment,
		ports,
		volumeBindings,
		healthChecks,
		svc.Networks,
		domainModel.ServiceTag(svc.Tag),
	), nil
}
