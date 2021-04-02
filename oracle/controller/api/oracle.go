package api

import (
	"context"
	"github.com/sirupsen/logrus"
	"oracle/service"
)

type Oracle struct {
	log     *logrus.Logger
	service *service.Service
	context context.Context
}

func NewOracle(ctx context.Context, log *logrus.Logger, service *service.Service) (*Oracle, error) {
	return &Oracle{log: log, context: ctx, service: service}, nil
}
