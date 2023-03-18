package logic

import (
	"context"

	"web/apps/property/rpc/internal/svc"
	"web/apps/property/rpc/property"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTableLogic {
	return &CreateTableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建表
func (l *CreateTableLogic) CreateTable(
	in *property.CreateTableReq,
) (*property.CreateTableRsp, error) {
	_ = l.svcCtx.UserScoreModel.InsertNew(l.ctx, in.Uid)

	return &property.CreateTableRsp{}, nil
}
