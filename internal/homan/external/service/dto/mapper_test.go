package dto

import (
	domainModel "github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapper(t *testing.T) {
	svc := domainModel.NewServiceConfig(
		"service-name",
		"domain.com",
		"image-name",
		[]string{"environment1"},
		[]domainModel.Port{
			domainModel.NewPort(5555),
			domainModel.NewPortBinding(443, 4433),
			domainModel.NewPortBindingTCP(80, 8080),
			domainModel.NewPortBindingUDP(90, 9090),
		},
		[]domainModel.Volume{
			domainModel.NewVolume("containerPath1"),
			domainModel.NewVolumeBinding("hostPath2", "containerPath2"),
			domainModel.NewVolumeBindingCopy("hostPath3", "containerPath3"),
		},
		[]domainModel.HealthCheck{
			domainModel.NewHealthCheckHTTP(120, "endpoint"),
			domainModel.NewHealthCheckTCP(121),
		},
		[]string{"network-name"},
		domainModel.TagGateway,
	)

	ext, err := MapServiceConfigToExternal(svc)
	assert.Nil(t, err)

	svcParse, err := MapExternalToServiceConfig("service-name", ext)
	assert.Nil(t, err)
	assert.NotEmpty(t, svcParse)

	assert.Equal(t, svcParse.IsValid(), svc.IsValid())
	assert.Equal(t, svcParse.Name(), svc.Name())
	assert.Equal(t, svcParse.DomainName(), svc.DomainName())
	assert.Equal(t, svcParse.Image(), svc.Image())
	assert.Len(t, svcParse.Environments(), len(svc.Environments()))
	for i, env := range svcParse.Environments() {
		assert.Equal(t, env, svc.Environments()[i])
	}
	assert.Len(t, svcParse.PortBindings(), len(svc.PortBindings()))
	for i, port := range svcParse.PortBindings() {
		assertPortBinding(t, port, svc.PortBindings()[i])
	}
	assert.Len(t, svcParse.VolumeBindings(), len(svc.VolumeBindings()))
	for i, volume := range svcParse.VolumeBindings() {
		assertVolumeBinding(t, volume, svc.VolumeBindings()[i])
	}
	assert.Len(t, svcParse.HealthChecks(), len(svc.HealthChecks()))
	for i, healthCheck := range svcParse.HealthChecks() {
		assertHealthCheck(t, healthCheck, svc.HealthChecks()[i])
	}
	assert.Len(t, svcParse.Networks(), len(svc.Networks()))
	for i, network := range svcParse.Networks() {
		assert.Equal(t, network, svc.Networks()[i])
	}
	assert.Equal(t, svcParse.Tag(), svc.Tag())
	assert.Equal(t, svcParse.IsCustom(), svc.IsCustom())
}

func assertPortBinding(t *testing.T, actual, expected domainModel.Port) {
	assert.NotEmpty(t, actual)
	assert.Equal(t, expected.IsValid(), actual.IsValid())
	assert.Equal(t, expected.ContainerPort(), actual.ContainerPort())
	assert.Equal(t, expected.HostPort(), actual.HostPort())
	assert.Equal(t, expected.Protocol(), actual.Protocol())
	assert.Equal(t, expected.String(), actual.String())
}

func assertVolumeBinding(t *testing.T, actual, expected domainModel.Volume) {
	assert.NotEmpty(t, actual)
	assert.Equal(t, expected.IsValid(), actual.IsValid())
	assert.Equal(t, expected.NeedCopy(), actual.NeedCopy())
	assert.Equal(t, expected.ContainerPath(), actual.ContainerPath())
	assert.Equal(t, expected.HostPath(), actual.HostPath())
	assert.Equal(t, expected.String(), actual.String())
}

func assertHealthCheck(t *testing.T, actual, expected domainModel.HealthCheck) {
	assert.NotEmpty(t, actual)
	assert.Equal(t, expected.IsValid(), actual.IsValid())
	assert.Equal(t, expected.Port(), actual.Port())
	assert.Equal(t, expected.Type(), actual.Type())
	assert.Equal(t, expected.Endpoint(), actual.Endpoint())
}
