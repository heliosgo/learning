package logic

import (
	"context"

	"web/apps/property/rpc/property"
	"web/apps/user/rpc/internal/svc"
	"web/apps/user/rpc/model"
	"web/apps/user/rpc/user"
	"web/pkg/errx"
	"web/pkg/tool"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	var res user.LoginResp
	u, err := l.svcCtx.UserModel.GetByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			/*
				return nil, errors.Wrapf(
					errx.NewWithCode(errx.DBError),
					"query user by username failed, username: %s, err: %v",
					in.Username,
					err,
				)
			*/
			u, err = l.register(in)
			if err != nil {
				return nil, err
			}

			_ = copier.Copy(&res, u)

			return &res, nil

		}
		return nil, err
	}

	password, _ := tool.Md5ByString(in.Password)
	if password != u.Password {
		return nil, errors.Wrapf(
			errx.NewWithCode(errx.RequestParamError),
			"login failed, password is incorrect, password: %s",
			in.Password,
		)
	}

	_ = copier.Copy(&res, u)

	return &res, nil
}

func (l *LoginLogic) register(in *user.LoginReq) (*model.User, error) {
	if len(in.Username) < 5 || len(in.Password) < 6 {
		return nil, errors.Wrapf(
			errx.NewWithCode(errx.RequestParamError),
			"username or password is not long enough, username: %s, password: %s",
			in.Username,
			in.Password,
		)
	}

	password, _ := tool.Md5ByString(in.Password)
	id, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: in.Username,
		Password: password,
	})
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"register failed, username: %s, password: %s, md5 password: %s",
			in.Username,
			in.Password,
			password,
		)
	}

	l.svcCtx.PropertyRPC.CreateTable(l.ctx, &property.CreateTableReq{
		Uid: id,
	})

	return &model.User{Id: id, Username: in.Username, Password: password}, nil
}
