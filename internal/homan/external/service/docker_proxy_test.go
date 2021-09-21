package service

import (
	"context"
	"fmt"
	"github.com/anantadwi13/homan/internal/homan/domain"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/stretchr/testify/assert"
	"net"
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

	proxyTCP, stopTCP, err := proxy.Start(context.TODO(), service.ProxyParams{
		Type:        model.ProxyTCP,
		TCPHostname: nginx.Name(),
		TCPPort:     80,
	})
	assert.Nil(t, err)
	assert.True(t, proxyTCP.IsRunning)
	defer stopTCP()

	proxyHTTP, stopHTTP, err := proxy.Start(context.TODO(), service.ProxyParams{Type: model.ProxyHTTP})
	assert.Nil(t, err)
	assert.True(t, proxyHTTP.IsRunning)
	defer stopHTTP()

	client := http.DefaultClient

	// Call nginx, should return OK
	request, err := http.NewRequest(http.MethodGet, proxyHTTP.FullScheme, nil)
	assert.Nil(t, err)
	request.Header.Add("X-Target-Host", "http://test-container")
	response, err := client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	request, err = http.NewRequest(http.MethodGet, fmt.Sprintf("http://%v/", proxyTCP.FullScheme), nil)
	assert.Nil(t, err)
	response, err = client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Stop nginx
	err = de2.Stop(context.TODO(), nginx)
	assert.Nil(t, err)

	// Call nginx, should return Not OK (Internal Server Error)
	request, err = http.NewRequest(http.MethodGet, proxyHTTP.FullScheme, nil)
	assert.Nil(t, err)
	request.Header.Add("X-Target-Host", "http://test-container")
	response, err = client.Do(request)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadGateway, response.StatusCode)

	dial, err := net.Dial("tcp", proxyTCP.Host)
	assert.Nil(t, err)
	assert.NotNil(t, dial)
	err = dial.Close()
	assert.Nil(t, err)

	_ = stopHTTP()
	_ = stopTCP()
	assert.False(t, proxyHTTP.IsRunning)
	assert.False(t, proxyTCP.IsRunning)
}
