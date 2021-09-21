package service

import (
	"context"
	"github.com/anantadwi13/homan/internal/homan/domain"
	"github.com/anantadwi13/homan/internal/homan/domain/model"
	"github.com/anantadwi13/homan/internal/homan/domain/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	r  service.Registry
	c  domain.Config
	de service.Executor
)

func init() {
	var err error
	cmd := NewCommander()
	c, err = domain.NewConfig(domain.ConfigParams{
		BasePath:    "/tmp/test-project-homan",
		ProjectName: "test-project",
		DaemonPort:  32321,
	})

	if err != nil {
		panic(err)
	}
	storage := service.NewStorage()
	r = NewLocalRegistry(c, storage)
	de = NewDockerExecutor(c, cmd, r, storage)
}

func TestDockerExecutor(t *testing.T) {
	sc := model.NewServiceConfig(
		"test-container",
		"test",
		"anantadwi13/docker-proxy",
		nil,
		[]model.Port{model.NewPortBinding(8081, 80)},
		nil,
		[]model.HealthCheck{model.NewHealthCheckTCP(80)},
		[]string{c.ProjectName()},
		model.TagWeb,
	)

	err := de.RunWait(context.TODO(), 10, r.GetCoreDaemon(context.TODO()))
	assert.Nil(t, err)

	defer func() {
		_ = de.Stop(context.TODO(), r.GetCoreDaemon(context.TODO()))
	}()

	err = de.RunWait(context.TODO(), 10, sc)
	assert.Nil(t, err)
	isRunning, err := de.IsRunning(context.TODO(), sc, false)
	assert.Nil(t, err)
	assert.Equal(t, true, isRunning)

	err = de.Restart(context.TODO(), sc)
	assert.Nil(t, err)
	isRunning, err = de.IsRunning(context.TODO(), sc, false)
	assert.Nil(t, err)
	assert.Equal(t, true, isRunning)

	err = de.Stop(context.TODO(), sc)
	assert.Nil(t, err)
	isRunning, err = de.IsRunning(context.TODO(), sc, false)
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
