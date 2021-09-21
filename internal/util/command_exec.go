package util

import (
	"context"
	"io"
	"os/exec"
)

func ExecCommand(ctx context.Context, name string, args ...string) ([]byte, error) {
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
