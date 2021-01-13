package service

import (
	"context"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture"
)

type Service struct {
	Context context.Context
}

func (service *Service) DB() *infrasturcture.Database {
	return infrasturcture.TxDB(service.Context)
}
