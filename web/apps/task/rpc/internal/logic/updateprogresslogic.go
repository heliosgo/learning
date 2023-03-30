package logic

import (
	"context"

	"web/apps/task/rpc/internal/svc"
	"web/apps/task/rpc/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProgressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProgressLogic {
	return &UpdateProgressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProgressLogic) UpdateProgress(in *task.UpdateProgressReq) (*task.UpdateProgressRsp, error) {
	// todo: add your logic here and delete this line

	return &task.UpdateProgressRsp{}, nil
}
