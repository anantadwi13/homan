package service

import (
	"context"
	"io"
	"os/exec"
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
	cmd := exec.CommandContext(ctx, name, args...)
	raw, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	resBytes, err := io.ReadAll(raw)
	if err != nil {
		return nil, err
	}
	errBytes, err := io.ReadAll(stdErr)
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return errBytes, err
	}
	return resBytes, nil
}
