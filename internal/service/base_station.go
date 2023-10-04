package service

import (
	"context"
	"ueransim-api/pkg/logger"
)

type BaseStationService struct {
	UeransimContainer string
	Executor          Executor
}

func NewBaseStationService(executor Executor) *BaseStationService {
	return &BaseStationService{
		UeransimContainer: "ueransim",
		Executor:          executor,
	}
}

func (s *BaseStationService) GetList(ctx context.Context) (string, error) {
	cmd := []string{"sh", "-c", "/ueransim/nr-cli -d"}

	result, err := s.Executor.Execute(ctx, s.UeransimContainer, cmd)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	logger.Infof("exec result: %+v", result)

	return result.StdOut, nil
}

func (s *BaseStationService) GetStatus(ctx context.Context, baseStation string) (string, error) {
	cmd := []string{"sh", "-c", "/ueransim/nr-cli", baseStation, "-e", "status"}

	result, err := s.Executor.Execute(ctx, s.UeransimContainer, cmd)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	logger.Infof("exec result: %+v", result)

	return result.StdOut, nil
}

func (s *BaseStationService) GetCountConnections(ctx context.Context, baseStation string) (string, error) {
	cmd := []string{"sh", "-c", "/ueransim/nr-cli", baseStation, "-e", "ue-count"}

	result, err := s.Executor.Execute(ctx, s.UeransimContainer, cmd)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	logger.Infof("exec result: %+v", result)

	return result.StdOut, nil
}

func (s *BaseStationService) GetInfo(ctx context.Context, baseStation string) (string, error) {
	cmd := []string{"sh", "-c", "/ueransim/nr-cli", baseStation, "-e", "ue-list"}

	result, err := s.Executor.Execute(ctx, s.UeransimContainer, cmd)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	logger.Infof("exec result: %+v", result)

	return result.StdOut, nil
}
