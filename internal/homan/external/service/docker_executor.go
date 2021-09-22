package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anantadwi13/homan/internal/homan/domain"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/anantadwi13/homan/internal/homan/external/api/homand"
	"github.com/anantadwi13/homan/internal/homan/external/service/dto"
	"net/http"
	"time"
)

type dockerExecutor struct {
	c            domain.Config
	cmd          Commander
	registry     service.Registry
	storage      service.Storage
	homandClient *homand.ClientWithResponses
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
	homandClient, err := homand.NewClientWithResponses(fmt.Sprintf("http://127.0.0.1:%v/", c.DaemonPort()))
	if err != nil {
		panic(err)
	}
	return &dockerExecutor{
		c:            c,
		cmd:          cmd,
		registry:     registry,
		storage:      storage,
		homandClient: homandClient,
	}
}

func (d dockerExecutor) Init(ctx context.Context, configs ...model.ServiceConfig) error {
	var needCopyConfigs []model.ServiceConfig

	for _, config := range configs {
		if config == nil {
			continue
		}
		for _, volume := range config.VolumeBindings() {
			if volume.NeedCopy() {
				needCopyConfigs = append(needCopyConfigs, config)
				break
			}
		}
	}

	err := d.copyVolume(ctx, needCopyConfigs...)
	if err != nil {
		return err
	}
	return nil
}

func (d *dockerExecutor) copyVolume(ctx context.Context, configs ...model.ServiceConfig) (err error) {
	stopped := false
	var newServices []model.ServiceConfig
	shutdown := func() {
		if stopped {
			return
		}
		stopped = true
		for _, config := range newServices {
			err2 := d.Stop(ctx, config)
			if err2 != nil {
				if err == nil {
					err = err2
				}
			}
			err2 = d.wait(ctx, config, false, 60, false)
			if err2 != nil {
				if err == nil {
					err = err2
				}
			}
		}
	}

	defer shutdown()

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
			config.Networks(),
			config.Tag(),
		)

		err := d.Run(ctx, newService)
		if err != nil {
			return err
		}

		err = d.wait(ctx, newService, true, 60, false)
		if err != nil {
			return err
		}

		newServices = append(newServices, newService)

		for _, volume := range config.VolumeBindings() {
			if !volume.NeedCopy() {
				continue
			}

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
	shutdown()
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

		if isRunning, err := d.IsRunning(ctx, config, false); err != nil || isRunning {
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

func (d *dockerExecutor) RunWait(ctx context.Context, timeout int, configs ...model.ServiceConfig) error {
	err := d.Run(ctx, configs...)
	if err != nil {
		return err
	}
	for _, config := range configs {
		if config == nil {
			continue
		}
		err := d.wait(ctx, config, true, timeout, true)
		if err != nil {
			_ = d.Stop(ctx, configs...)
			return err
		}
	}
	return nil
}

func (d *dockerExecutor) wait(
	ctx context.Context, config model.ServiceConfig, expectRunning bool, timeout int, checkHealth bool,
) error {
	isRunning, err := d.IsRunning(ctx, config, checkHealth)
	timeoutPerIter := 1000 // in ms
	retry := timeout * 1000 / timeoutPerIter
	for retry > 0 && (err == nil && isRunning != expectRunning) {
		time.Sleep(time.Duration(timeoutPerIter) * time.Millisecond)
		isRunning, err = d.IsRunning(ctx, config, checkHealth)
		retry--
	}
	if err != nil {
		return err
	}
	if isRunning != expectRunning {
		return errors.New(fmt.Sprintf("waiting of %v is reaching timeout", config.Name()))
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
		if isRunning, err := d.IsRunning(ctx, config, false); err != nil || !isRunning {
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

func (d *dockerExecutor) IsRunning(ctx context.Context, config model.ServiceConfig, checkHealth bool) (
	bool, error,
) {
	isRunning, err := d.isRunning(ctx, config)
	if err != nil {
		return false, err
	}
	if !isRunning {
		return false, nil
	}

	if checkHealth {
		isRunning, err := d.isRunning(ctx, d.registry.GetCoreDaemon(ctx))
		if err != nil {
			return false, err
		}
		if !isRunning {
			return false, errors.New("core daemon is not running")
		}

		for _, healthCheck := range config.HealthChecks() {
			if healthCheck == nil {
				continue
			}

			switch {
			case config.Name() == d.registry.GetCoreDaemon(ctx).Name():
				_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%v/", d.c.DaemonPort()))
				if err != nil {
					return false, nil
				}
				return true, nil
			default:
				var (
					hcType  homand.CheckHealthJSONBodyCheckType
					address string
				)

				switch healthCheck.Type() {
				case model.HealthCheckHTTP:
					hcType = "http"
					address = fmt.Sprintf("http://%v:%v%v", config.Name(), healthCheck.Port(), healthCheck.Endpoint())
				case model.HealthCheckTCP:
					hcType = "tcp"
					address = fmt.Sprintf("%v:%v", config.Name(), healthCheck.Port())
				default:
					continue
				}

				res, err := d.homandClient.CheckHealthWithResponse(ctx, homand.CheckHealthJSONRequestBody{
					Address:   address,
					CheckType: hcType,
				})
				if err != nil {
					return false, err
				}
				if res.JSON200 == nil || !res.JSON200.IsAvailable {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

func (d *dockerExecutor) isRunning(ctx context.Context, config model.ServiceConfig) (bool, error) {
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
	if containers[0] == nil || !containers[0].State.Running {
		return false, nil
	}
	return true, nil
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
		isRunning, err := d.IsRunning(ctx, config, false)
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
