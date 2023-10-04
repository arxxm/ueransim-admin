package service

import (
	"context"
)

type Service struct {
	BaseStation    BaseStation
	MobileTerminal MobileTerminal
}

func NewService(baseStation BaseStation, mobileTerminal MobileTerminal) *Service {
	return &Service{
		BaseStation:    baseStation,
		MobileTerminal: mobileTerminal,
	}
}

type BaseStation interface {
	GetList(ctx context.Context) (string, error)
	GetStatus(ctx context.Context, baseStation string) (string, error)
	GetCountConnections(ctx context.Context, baseStation string) (string, error)
	GetInfo(ctx context.Context, baseStation string) (string, error)
}

type MobileTerminal interface {
	GetList(ctx context.Context) (string, error)
	GetStatus(ctx context.Context, mobileTerminal string) (string, error)
	GetNetworkConnectionStatus(ctx context.Context, mobileTerminal string) (string, error)
}

type Executor interface {
	Execute(ctx context.Context, container string, cmd []string) (ExecResult, error)
}
