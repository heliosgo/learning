// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"web/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq  = user.LoginReq
	LoginResp = user.LoginResp
	Request   = user.Request
	Response  = user.Response
	UserInfo  = user.UserInfo

	User interface {
		// 登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// 登录
func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
