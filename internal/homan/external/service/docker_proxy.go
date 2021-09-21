package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/anantadwi13/homan/internal/homan/domain"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/google/uuid"
	"net"
	"strconv"
	"time"
)

type dockerProxy struct {
	config   domain.Config
	executor service.Executor
}

func NewDockerProxy(config domain.Config, executor service.Executor) service.Proxy {
	return &dockerProxy{config: config, executor: executor}
}

func (d *dockerProxy) Start(ctx context.Context, params service.ProxyParams) (
	proxy *model.ProxyDetail, stop func() error, err error,
) {
	proxyLocalPort := d.proxyPort()
	proxyContainerPort := 5555
	var env []string

	switch params.Type {
	case model.ProxyTCP:
		if params.TCPHostname == "" || params.TCPPort < 0 || params.TCPPort > 65535 {
			err = errors.New("invalid tcp proxy parameters (hostname or port)")
			return
		}

		env = []string{
			"DOCKER_PROXY_MODE=tcp",
			fmt.Sprintf("DOCKER_PROXY_LOCAL_ADDRESS=:%v", proxyContainerPort),
			fmt.Sprintf("DOCKER_PROXY_REMOTE_ADDRESS=%v:%v", params.TCPHostname, params.TCPPort),
		}
		proxy = &model.ProxyDetail{
			Type:       model.ProxyTCP,
			Host:       fmt.Sprintf("127.0.0.1:%v", proxyLocalPort),
			FullScheme: fmt.Sprintf("127.0.0.1:%v", proxyLocalPort),
		}
	case model.ProxyHTTP:
		env = []string{
			"DOCKER_PROXY_MODE=http",
			fmt.Sprintf("DOCKER_PROXY_LOCAL_ADDRESS=:%v", proxyContainerPort),
		}
		proxy = &model.ProxyDetail{
			Type:       model.ProxyHTTP,
			Host:       fmt.Sprintf("%v:%v", "127.0.0.1", proxyLocalPort),
			FullScheme: fmt.Sprintf("http://%v:%v/", "127.0.0.1", proxyLocalPort),
		}
	default:
		err = errors.New("unknown proxy type")
		return
	}

	proxyService := model.NewServiceConfig(
		d.proxyName(),
		"",
		"anantadwi13/docker-proxy:0.2.0",
		env,
		[]model.Port{model.NewPortBinding(proxyLocalPort, proxyContainerPort)},
		[]model.Volume{},
		[]model.HealthCheck{model.NewHealthCheckTCP(proxyContainerPort)},
		[]string{d.config.ProjectName()},
		model.TagProxy,
	)

	// Start Proxy
	err = d.executor.RunWait(ctx, 10, proxyService)
	if err != nil {
		return
	}
	proxy.IsRunning = true

	stop = func() error {
		proxy.IsRunning = false
		err := d.executor.Stop(ctx, proxyService)
		if err != nil {
			return err
		}
		return nil
	}

	return
}

func (d *dockerProxy) proxyName() string {
	return d.config.SystemNamePrefix() + "proxy-" + uuid.New().String()
}

func (d *dockerProxy) proxyPort() int {
	selectedPort := -1
	for port := 20001; port < 65536; port++ {
		timeout := 10 * time.Millisecond
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("127.0.0.1", strconv.Itoa(port)), timeout)
		if err != nil || conn == nil {
			selectedPort = port
			break
		}
		_ = conn.Close()
	}
	return selectedPort
}
