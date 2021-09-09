package service

import (
	"context"
	"fmt"
	"github.com/anantadwi13/cli-whm/internal/domain"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
)

type dockerProxy struct {
	config   domain.Config
	executor service.Executor
}

func NewDockerProxy(config domain.Config, executor service.Executor) service.Proxy {
	return &dockerProxy{config, executor}
}

func (d *dockerProxy) Execute(ctx context.Context, request func(proxy *model.ProxyDetail) error) (err error) {
	proxyService := model.NewServiceConfig(
		d.proxyName(),
		"",
		"anantadwi13/docker-proxy:0.1.0",
		[]string{},
		[]model.Port{model.NewPortBinding(d.proxyPort(), 80)},
		[]model.Volume{},
		[]string{d.config.ProjectName()},
		model.TagProxy,
	)
	proxyDetail := &model.ProxyDetail{Host: fmt.Sprintf("http://%v:%v", "localhost", d.proxyPort())}

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

	// Send Request Through Proxy
	err = request(proxyDetail)
	if err != nil {
		return
	}

	return nil
}

func (d *dockerProxy) proxyName() string {
	return d.config.SystemNamePrefix() + "proxy"
}

func (d *dockerProxy) proxyPort() int {
	return 5555
}
