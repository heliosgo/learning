package user

import (
	"context"
	"time"

	"web/apps/app/internal/svc"
	"web/apps/app/internal/types"
	"web/apps/user/rpc/user"
	"web/pkg/jwtx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	var loginReq user.LoginReq
	loginReq.Username = req.Username
	loginReq.Password = req.Password
	res, err := l.svcCtx.UserRPC.Login(l.ctx, &loginReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := jwtx.GetToken(
		l.svcCtx.Config.JwtAuth.AccessSecret,
		now,
		accessExpire,
		res.Id,
	)

	return &types.LoginResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}
