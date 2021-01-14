// 服务层
// 由于采用了OpenSessionInView的设计，事务以Controller为边界
// 因此对于没有额外业务的数据库CRUD来说，服务层并不是必须的，可以在Controller中直接调用模型的方法
// 但是有较多操作，或者有明确重用需求时，仍建议封装在Service层中
package service

import (
	"context"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
)

type Service struct {
	Context context.Context
}

func (service *Service) DB() *db.Database {
	return db.TxDB(service.Context)
}
