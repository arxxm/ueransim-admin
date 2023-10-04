package service

import (
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io/ioutil"
	"ueransim-api/pkg/logger"
)

type ExecutorService struct{}

func NewExecutorService() *ExecutorService {
	return &ExecutorService{}
}

type ExecResult struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

func (e ExecutorService) Execute(ctx context.Context, container string, cmd []string) (ExecResult, error) {
	ex := ExecResult{}

	respID, err := e.exec(ctx, container, cmd)
	if err != nil {
		logger.Error(err)
		return ex, err
	}

	ex, err = e.inspectExecResp(ctx, respID.ID)
	if err != nil {
		logger.Error(err)
		return ex, err
	}

	return ex, nil
}

func (e ExecutorService) exec(ctx context.Context, containerID string, command []string) (types.IDResponse, error) {
	docker, err := client.NewEnvClient()
	if err != nil {
		return types.IDResponse{}, err
	}
	defer docker.Close()

	config := types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          command,
	}

	return docker.ContainerExecCreate(ctx, containerID, config)
}

func (e ExecutorService) inspectExecResp(ctx context.Context, id string) (ExecResult, error) {
	var execResult ExecResult
	docker, err := client.NewEnvClient()
	if err != nil {
		return execResult, err
	}
	defer docker.Close()

	resp, err := docker.ContainerExecAttach(ctx, id, types.ExecStartCheck{})
	if err != nil {
		return execResult, err
	}
	defer resp.Close()

	// read the output
	var outBuf, errBuf bytes.Buffer
	outputDone := make(chan error)

	go func() {
		_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
		outputDone <- err
	}()

	select {
	case err := <-outputDone:
		if err != nil {
			return execResult, err
		}
		break

	case <-ctx.Done():
		return execResult, ctx.Err()
	}

	stdout, err := ioutil.ReadAll(&outBuf)
	if err != nil {
		return execResult, err
	}
	stderr, err := ioutil.ReadAll(&errBuf)
	if err != nil {
		return execResult, err
	}

	res, err := docker.ContainerExecInspect(ctx, id)
	if err != nil {
		return execResult, err
	}

	execResult.ExitCode = res.ExitCode
	execResult.StdOut = string(stdout)
	execResult.StdErr = string(stderr)
	return execResult, nil
}
