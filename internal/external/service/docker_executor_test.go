package service

import (
	"context"
	"github.com/anantadwi13/cli-whm/internal/domain"
	"github.com/anantadwi13/cli-whm/internal/domain/model"
	"github.com/anantadwi13/cli-whm/internal/domain/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	r  service.Registry
	de service.Executor
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
		[]string{"test-network"},
		model.TagWeb,
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
