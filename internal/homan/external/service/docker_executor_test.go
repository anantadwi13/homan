package service

import (
	"context"
	"github.com/anantadwi13/homan/internal/homan/domain"
	model2 "github.com/anantadwi13/homan/internal/homan/domain/model"
	service2 "github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	r  service2.Registry
	de service2.Executor
)

func init() {
	cmd := NewCommander()
	c, err := domain.NewConfig(domain.ConfigParams{
		BasePath:    "../../../temp",
		ProjectName: "test-project",
	})

	if err != nil {
		panic(err)
	}
	storage := service2.NewStorage()
	r = NewLocalRegistry(c, storage)
	de = NewDockerExecutor(c, cmd, r, storage)
}

func TestDockerExecutor(t *testing.T) {
	sc := model2.NewServiceConfig(
		"test-container",
		"test",
		"anantadwi13/docker-proxy",
		nil,
		[]model2.Port{model2.NewPortBinding(8081, 80)},
		nil,
		[]model2.HealthCheck{model2.NewHealthCheckHTTP(80, "/")},
		[]string{"test-network"},
		model2.TagWeb,
	)

	err := de.Run(context.TODO(), sc)
	assert.Nil(t, err)
	isRunning, err := de.IsRunning(context.TODO(), sc)
	assert.Nil(t, err)
	assert.Equal(t, true, isRunning)

	err = de.Restart(context.TODO(), sc)
	assert.Nil(t, err)
	isRunning, err = de.IsRunning(context.TODO(), sc)
	assert.Nil(t, err)
	assert.Equal(t, true, isRunning)

	err = de.Stop(context.TODO(), sc)
	assert.Nil(t, err)
	isRunning, err = de.IsRunning(context.TODO(), sc)
	assert.Nil(t, err)
	assert.Equal(t, false, isRunning)

	err = r.Add(context.TODO(), sc)
	assert.Nil(t, err)

	err = r.Remove(context.TODO(), sc)
	assert.Nil(t, err)
	all, err := r.GetAll(context.TODO())
	assert.Nil(t, err)
	assert.Len(t, all, 0)
}
