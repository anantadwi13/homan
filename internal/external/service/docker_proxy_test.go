package service

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/domain"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	proxy service.Proxy
	c2    domain.Config
	de2   service.Executor
)

func init() {
	cmd := NewCommander()
	var err error
	c2, err = domain.NewConfig(domain.ConfigParams{
		BasePath:    "../../../temp",
		ProjectName: "test-project",
	})

	if err != nil {
		panic(err)
	}
	storage := service.NewStorage()
	r := NewLocalRegistry(c2, storage)
	de2 = NewDockerExecutor(c2, cmd, r, storage)
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
		[]string{c2.ProjectName()},
		model.TagWeb,
	)

	err := de2.Run(context.TODO(), nginx)
	assert.Nil(t, err)

	err = proxy.Execute(context.TODO(), func(proxy *model.ProxyDetail) error {
		client := http.DefaultClient

		// Call nginx, should return OK
		request, err := http.NewRequest(http.MethodGet, proxy.Host, nil)
		assert.Nil(t, err)
		request.Header.Add("X-Target-Host", "http://test-container")
		response, err := client.Do(request)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		err = de2.Stop(context.TODO(), nginx)
		assert.Nil(t, err)

		// Call nginx, should return Not OK (Internal Server Error)
		request, err = http.NewRequest(http.MethodGet, proxy.Host, nil)
		assert.Nil(t, err)
		request.Header.Add("X-Target-Host", "http://test-container")
		response, err = client.Do(request)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

		return nil
	})
	assert.Nil(t, err)
}
