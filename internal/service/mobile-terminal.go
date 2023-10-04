package service

import (
	"context"
	"ueransim-api/pkg/logger"
)

type MobileTerminalService struct {
	UeransimContainer string
	Executor          Executor
}

func NewMobileTerminalService(executor Executor) *MobileTerminalService {
	return &MobileTerminalService{
		UeransimContainer: "ueransim",
		Executor:          executor,
	}
}

func (s *MobileTerminalService) GetList(ctx context.Context) (string, error) {
	cmd := []string{"sh", "-c", "/ueransim/nr-cli -d"}

	result, err := s.Executor.Execute(ctx, s.UeransimContainer, cmd)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	logger.Infof("exec result: %+v", result)

	return result.StdOut, nil
}

func (s *MobileTerminalService) GetStatus(ctx context.Context, mobileTerminal string) (string, error) {
	cmd := []string{"sh", "-c", "/ueransim/nr-cli", mobileTerminal, "-e", "status"}

	result, err := s.Executor.Execute(ctx, s.UeransimContainer, cmd)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	logger.Infof("exec result: %+v", result)

	return result.StdOut, nil
}

func (s *MobileTerminalService) GetNetworkConnectionStatus(ctx context.Context, mobileTerminal string) (string, error) {
	cmd := []string{"sh", "-c", "/ueransim/nr-cli", mobileTerminal, "-e", "ps-list"}

	result, err := s.Executor.Execute(ctx, s.UeransimContainer, cmd)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	logger.Infof("exec result: %+v", result)

	return result.StdOut, nil
}
