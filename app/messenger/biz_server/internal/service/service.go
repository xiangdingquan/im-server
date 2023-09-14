package service

import (
	"context"
)

type Service struct {
}

func New() *Service {
	return new(Service)
}

func (s *Service) Close() error {
	return nil
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return nil
}
