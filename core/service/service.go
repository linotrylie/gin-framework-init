package service

import (
	"context"
	"equity/core/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	service := Service{ctx: ctx}
	return service
}
