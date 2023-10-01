package service

import (
	"context"
)

type Service struct {
	EmulatorObjects EmulatorObjects
}

func NewService(emulatorObjects EmulatorObjects) *Service {
	return &Service{
		EmulatorObjects: emulatorObjects,
	}
}

type EmulatorObjects interface {
	GetList(ctx context.Context) (string, error)
}
