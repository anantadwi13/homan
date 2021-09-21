package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anantadwi13/cli-whm/internal/homan/domain"
	"github.com/anantadwi13/cli-whm/internal/homan/domain/model"
	"github.com/anantadwi13/cli-whm/internal/homan/domain/service"
	"github.com/anantadwi13/cli-whm/internal/homan/external/service/dto"
	"time"
)

type dockerExecutor struct {
	c        domain.Config
	cmd      Commander
	registry service.Registry
	storage  service.Storage
}

var ErrorDockerExecutorNetworkExist = errors.New("error [docker_executor]: network already exist")
var ErrorDockerExecutorNetworkNotExist = errors.New("error [docker_executor]: network does not exist")
var ErrorDockerExecutorNetworkBeingUsed = errors.New("error [docker_executor]: network is being used")

func NewDockerExecutor(
	c domain.Config,
	cmd Commander,
	registry service.Registry,
	storage service.Storage,
) service.Executor {
	return &dockerExecutor{
		c:        c,
		cmd:      cmd,
		registry: registry,
		storage:  storage,
	}
}

func (d *dockerExecutor) InitVolume(ctx context.Context, configs ...model.ServiceConfig) error {
	var newServices []model.ServiceConfig
	defer func() {
		for _, config := range newServices {
			_ = d.Stop(ctx, config)
		}
	}()

	for _, config := range configs {
		if config == nil {
			continue
		}
		if !config.IsValid() {
			return service.ErrorExecutorServiceConfigInvalid
		}

		newService := model.NewServiceConfig(
			config.Name(),
			config.DomainName(),
			config.Image(),
			config.Environments(),
			nil,
			nil,
			config.HealthChecks(),
			nil,
			config.Tag(),
		)

		err := d.Run(ctx, newService)
		if err != nil {
			return err
		}

		isRunning, err := d.IsRunning(ctx, newService)
		retry := 10
		for retry > 0 && (err == nil && !isRunning) {
			time.Sleep(100 * time.Millisecond)
			isRunning, err = d.IsRunning(ctx, newService)
			retry--
		}
		if err != nil {
			return err
		}

		newServices = append(newServices, newService)

		for _, volume := range config.VolumeBindings() {
			err := d.storage.Mkdir(volume.HostPath())
			if err != nil {
				return err
			}

			cmdRes, err := d.cmd.RunCommand(
				ctx,
				"docker",
				"cp",
				fmt.Sprintf("%v:%v/.", d.getContainerName(ctx, config), volume.ContainerPath()),
				volume.HostPath(),
			)
			if err != nil {
				return errors.New(string(cmdRes))
			}
		}
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
			return service.ErrorExecutorServiceConfigInvalid
		}

		if isRunning, err := d.IsRunning(ctx, config); err != nil || isRunning {
			if err == nil {
				err = service.ErrorExecutorServiceIsRunning
			}
			return err
		}

		args := []string{"run", "--rm", "-dit"}

		for _, env := range config.Environments() {
			args = append(args, "-e")
			args = append(args, env)
		}

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
			args = append(args, config.Name())
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
	var networks []string

	for _, config := range configs {
		if config == nil {
			continue
		}
		if !config.IsValid() {
			return service.ErrorExecutorServiceConfigInvalid
		}
		if isRunning, err := d.IsRunning(ctx, config); err != nil || !isRunning {
			if err == nil {
				err = service.ErrorExecutorServiceIsNotRunning
			}
			return err
		}
		serviceName = append(serviceName, d.getContainerName(ctx, config))
		networks = append(networks, config.Networks()...)
	}
	for _, containerName := range serviceName {
		cmdRes, err := d.cmd.RunCommand(ctx, "docker", "stop", containerName)
		if err != nil {
			return errors.New(string(cmdRes))
		}
	}
	for _, network := range networks {
		_ = d.removeNetwork(ctx, network)
	}
	return nil
}

func (d *dockerExecutor) IsRunning(ctx context.Context, config model.ServiceConfig) (bool, error) {
	if config == nil || !config.IsValid() {
		return false, service.ErrorExecutorServiceConfigInvalid
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
			return service.ErrorExecutorServiceConfigInvalid
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

func (d *dockerExecutor) removeNetwork(ctx context.Context, networkName string) error {
	if !d.isNetworkExist(ctx, networkName) {
		return ErrorDockerExecutorNetworkNotExist
	}

	if isUsed, err := d.isNetworkUsed(ctx, networkName); err != nil || isUsed {
		if err == nil {
			err = ErrorDockerExecutorNetworkBeingUsed
		}
		return err
	}

	resCmd, err := d.cmd.RunCommand(ctx, "docker", "network", "rm", networkName)
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

func (d *dockerExecutor) isNetworkUsed(ctx context.Context, networkName string) (bool, error) {
	cmdRes, err := d.cmd.RunCommand(ctx, "docker", "network", "inspect", networkName)
	if err != nil {
		return false, nil
	}
	var networks []*dto.DockerNetworkInspect
	err = json.Unmarshal(cmdRes, &networks)
	if err != nil {
		return false, err
	}
	if len(networks) == 0 || len(networks[0].Containers) <= 0 {
		return false, nil
	}
	return true, nil
}

func (d *dockerExecutor) getContainerName(ctx context.Context, config model.ServiceConfig) string {
	if config == nil {
		return ""
	}
	hostname := config.Name()
	return d.c.ProjectName() + "_" + hostname
}
