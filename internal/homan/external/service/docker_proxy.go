package service

import (
	"context"
	"fmt"
	"github.com/anantadwi13/cli-whm/internal/homan/domain"
	model2 "github.com/anantadwi13/cli-whm/internal/homan/domain/model"
	service2 "github.com/anantadwi13/cli-whm/internal/homan/domain/service"
	"github.com/google/uuid"
	"net"
	"net/http"
	"strconv"
	"time"
)

type dockerProxy struct {
	config   domain.Config
	executor service2.Executor
}

func NewDockerProxy(config domain.Config, executor service2.Executor) service2.Proxy {
	return &dockerProxy{config: config, executor: executor}
}

func (d *dockerProxy) Execute(ctx context.Context, request func(proxy *model2.ProxyDetail) error) (err error) {
	proxyService := model2.NewServiceConfig(
		d.proxyName(),
		"",
		"anantadwi13/docker-proxy:0.2.0",
		[]string{},
		[]model2.Port{model2.NewPortBinding(d.proxyPort(), 80)},
		[]model2.Volume{},
		// todo add healtcheck based on proxy type (tcp or http)
		[]model2.HealthCheck{},
		[]string{d.config.ProjectName()},
		model2.TagProxy,
	)
	proxyDetail := &model2.ProxyDetail{
		Host:     fmt.Sprintf("%v:%v", "localhost", d.proxyPort()),
		FullPath: fmt.Sprintf("http://%v:%v/", "localhost", d.proxyPort()),
	}

	// Start Proxy
	err = d.executor.Run(ctx, proxyService)
	if err != nil {
		return
	}

	defer func() {
		// Stop Proxy
		err2 := d.executor.Stop(ctx, proxyService)
		if err == nil {
			err = err2
		}
	}()

	// Wait until proxy is running
	c := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, proxyDetail.FullPath, nil)
	if err != nil {
		return err
	}
	do, err := c.Do(req)
	retry := 10
	for retry > 0 && (err != nil || do.StatusCode != http.StatusInternalServerError) {
		time.Sleep(100 * time.Millisecond)
		do, err = c.Do(req)
		retry--
	}
	if err != nil {
		return
	}

	// Send Request Through Proxy
	err = request(proxyDetail)
	if err != nil {
		return
	}

	return nil
}

func (d *dockerProxy) proxyName() string {
	return d.config.SystemNamePrefix() + "proxy-" + uuid.New().String()
}

func (d *dockerProxy) proxyPort() int {
	selectedPort := -1
	for port := 20001; port < 65536; port++ {
		timeout := 10 * time.Millisecond
		conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", strconv.Itoa(port)), timeout)
		if err != nil || conn == nil {
			selectedPort = port
			break
		}
		_ = conn.Close()
	}
	return selectedPort
}
