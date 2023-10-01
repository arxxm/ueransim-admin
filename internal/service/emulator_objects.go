package service

import (
	"context"
	"ueransim-api/pkg/logger"
)

type EmulatorObjectsService struct {
	UeransimContainer string
}

func NewEmulatorObjectsService() *EmulatorObjectsService {
	return &EmulatorObjectsService{
		UeransimContainer: "ueransim",
	}
}

func (s *EmulatorObjectsService) GetList(ctx context.Context) (string, error) {
	cmd := []string{"sh", "-c", "/ueransim/nr-cli -d"}

	respID, err := Exec(ctx, s.UeransimContainer, cmd)
	if err != nil {
		logger.Error("EmulatorObjectsService GetList err:", err)
	}

	ex, err := InspectExecResp(ctx, respID.ID)
	if err != nil {
		logger.Error("EmulatorObjectsService InspectExecResp err:", err)
	}

	logger.Infof("exec result: %+v", ex)

	return ex.StdOut, nil
}
