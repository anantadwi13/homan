package service

import (
	"context"
	"github.com/anantadwi13/homan/internal/homan/domain"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	proxy service.Proxy
	r2    service.Registry
	c2    domain.Config
	de2   service.Executor
)

func init() {
	cmd := NewCommander()
	var err error
	c2, err = domain.NewConfig(domain.ConfigParams{
		BasePath:    "/tmp/test-project-homan",
		ProjectName: "test-project",
		DaemonPort:  32321,
	})

	if err != nil {
		panic(err)
	}
	storage := service.NewStorage()
	r2 = NewLocalRegistry(c2, storage)
	de2 = NewDockerExecutor(c2, cmd, r2, storage)
	proxy = NewDockerProxy(c2, de2)
}

func TestDockerProxy(t *testing.T) {
	nginx := model.NewServiceConfig(
		"test-container",
		"test",
		"nginx",
		nil,
		[]model.Port{model.NewPort(80)},
		nil,
		[]model.HealthCheck{model.NewHealthCheckHTTP(80, "/")},
		[]string{c2.ProjectName()},
		model.TagWeb,
	)

	err := de.RunWait(context.TODO(), 10, r2.GetCoreDaemon(context.TODO()))
	assert.Nil(t, err)

	err = de2.RunWait(context.TODO(), 10, nginx)
	assert.Nil(t, err)

	defer func() {
		_ = de2.Stop(context.TODO(), nginx)
		_ = de2.Stop(context.TODO(), r2.GetCoreDaemon(context.TODO()))
	}()

	err = proxy.Execute(context.TODO(), func(proxy *model.ProxyDetail) error {
		client := http.DefaultClient

		// Call nginx, should return OK
		request, err := http.NewRequest(http.MethodGet, proxy.FullPath, nil)
		assert.Nil(t, err)
		request.Header.Add("X-Target-Host", "http://test-container")
		response, err := client.Do(request)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		err = de2.Stop(context.TODO(), nginx)
		assert.Nil(t, err)

		// Call nginx, should return Not OK (Internal Server Error)
		request, err = http.NewRequest(http.MethodGet, proxy.FullPath, nil)
		assert.Nil(t, err)
		request.Header.Add("X-Target-Host", "http://test-container")
		response, err = client.Do(request)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadGateway, response.StatusCode)

		return nil
	})
	assert.Nil(t, err)
}
