package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	domainService "github.com/anantadwi13/cli-whm/internal/domain/service"
	"github.com/anantadwi13/cli-whm/internal/external/service/dto"
)

type dockerExecutor struct {
	c        domainService.Config
	cmd      Commander
	registry domainService.Registry
}

var ErrorDockerExecutorNetworkExist = errors.New("error [docker_executor]: network already exist")

func NewDockerExecutor(
	c domainService.Config,
	cmd Commander,
	registry domainService.Registry,
) domainService.Executor {
	return &dockerExecutor{
		c:        c,
		cmd:      cmd,
		registry: registry,
	}
}

func (d *dockerExecutor) RunAll(ctx context.Context) error {
	all, err := d.registry.GetAll(ctx)
	if err != nil {
		return err
	}
	err = d.Run(ctx, all...)
	if err != nil {
		return err
	}
	return nil
}

func (d *dockerExecutor) Run(ctx context.Context, configs ...model.ServiceConfig) error {
	var serviceConfigs []model.ServiceConfig
	var configArgs [][]string
	var networks []string

	for _, config := range configs {
		if config == nil {
			continue
		}
		if !config.IsValid() {
			return domainService.ErrorExecutorServiceConfigInvalid
		}

		if isRunning, err := d.IsRunning(ctx, config); err != nil || isRunning {
			if err == nil {
				err = domainService.ErrorExecutorServiceIsRunning
			}
			return err
		}

		args := []string{"run", "--rm", "-dit"}

		for _, port := range config.PortBindings() {
			args = append(args, "-p")
			args = append(args, port.String())
		}

		for _, vol := range config.VolumeBindings() {
			args = append(args, "-v")
			args = append(args, vol.String())
		}

		args = append(args, "--name")
		args = append(args, d.getContainerName(ctx, config))

		for _, network := range config.Networks() {
			networks = append(networks, network)
			args = append(args, "--network")
			args = append(args, network)
			args = append(args, "--network-alias")
			args = append(args, d.getContainerHostname(ctx, config))
		}

		args = append(args, config.Image())

		configArgs = append(configArgs, args)
		serviceConfigs = append(serviceConfigs, config)
	}
	for _, network := range networks {
		if !d.isNetworkExist(ctx, network) {
			err := d.addNetwork(ctx, network)
			if err != nil {
				return err
			}
		}
	}

	for _, args := range configArgs {
		resCmd, err := d.cmd.RunCommand(ctx, "docker", args...)
		if err != nil {
			return errors.New(string(resCmd))
		}
	}
	return nil
}

func (d *dockerExecutor) Stop(ctx context.Context, configs ...model.ServiceConfig) error {
	var serviceName []string

	for _, config := range configs {
		if config == nil {
			continue
		}
		if !config.IsValid() {
			return domainService.ErrorExecutorServiceConfigInvalid
		}
		if isRunning, err := d.IsRunning(ctx, config); err != nil || !isRunning {
			if err == nil {
				err = domainService.ErrorExecutorServiceIsNotRunning
			}
			return err
		}
		serviceName = append(serviceName, d.getContainerName(ctx, config))
	}
	for _, containerName := range serviceName {
		cmdRes, err := d.cmd.RunCommand(ctx, "docker", "stop", containerName)
		if err != nil {
			return errors.New(string(cmdRes))
		}
	}
	return nil
}

func (d *dockerExecutor) IsRunning(ctx context.Context, config model.ServiceConfig) (bool, error) {
	if config == nil || !config.IsValid() {
		return false, domainService.ErrorExecutorServiceConfigInvalid
	}
	cmdRes, err := d.cmd.RunCommand(ctx, "docker", "inspect", d.getContainerName(ctx, config))
	if err != nil {
		// docker inspect returns error. it indicates that service is not running
		return false, nil
	}
	var containers []*dto.DockerContainerInspect
	err = json.Unmarshal(cmdRes, &containers)
	if err != nil {
		return false, err
	}
	if len(containers) == 0 {
		return false, nil
	}
	if containers[0] != nil && containers[0].State.Running {
		return true, nil
	}
	return false, nil
}

func (d *dockerExecutor) Restart(ctx context.Context, configs ...model.ServiceConfig) error {
	var servicesNeedToStop []model.ServiceConfig
	var servicesNeedToRun []model.ServiceConfig

	for _, config := range configs {
		if config == nil {
			continue
		}
		if !config.IsValid() {
			return domainService.ErrorExecutorServiceConfigInvalid
		}
		isRunning, err := d.IsRunning(ctx, config)
		if err != nil {
			return err
		}
		if isRunning {
			servicesNeedToStop = append(servicesNeedToStop, config)
		}
		servicesNeedToRun = append(servicesNeedToRun, config)
	}

	for _, config := range servicesNeedToStop {
		err := d.Stop(ctx, config)
		if err != nil {
			return err
		}
	}
	for _, config := range servicesNeedToRun {
		err := d.Run(ctx, config)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *dockerExecutor) addNetwork(ctx context.Context, networkName string) error {
	if d.isNetworkExist(ctx, networkName) {
		return ErrorDockerExecutorNetworkExist
	}

	resCmd, err := d.cmd.RunCommand(ctx, "docker", "network", "create", networkName)
	if err != nil {
		return errors.New(string(resCmd))
	}
	return nil
}
func (d *dockerExecutor) isNetworkExist(ctx context.Context, networkName string) bool {
	_, err := d.cmd.RunCommand(ctx, "docker", "network", "inspect", networkName)
	if err != nil {
		return false
	}
	return true
}

func (d *dockerExecutor) getContainerName(ctx context.Context, config model.ServiceConfig) string {
	if config == nil {
		return ""
	}
	hostname := d.getContainerHostname(ctx, config)
	return d.c.ProjectName() + "_" + hostname
}

func (d *dockerExecutor) getContainerHostname(ctx context.Context, config model.ServiceConfig) string {
	if config == nil {
		return ""
	}
	name := config.Name()
	if isSystem, err := d.registry.IsSystem(ctx, config); err == nil && isSystem {
		name = d.c.SystemNamePrefix() + config.Name()
	}
	return name
}
