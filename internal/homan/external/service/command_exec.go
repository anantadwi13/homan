package service

import (
	"context"
	"github.com/anantadwi13/homan/internal/util"
)

type Commander interface {
	RunCommand(ctx context.Context, name string, args ...string) (res []byte, err error)
}
type commander struct {
}

func NewCommander() Commander {
	return &commander{}
}

func (c commander) RunCommand(ctx context.Context, name string, args ...string) ([]byte, error) {
	return util.ExecCommand(ctx, name, args...)
}
