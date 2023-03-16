package model

import "context"

type UserModel interface {
	userModelReader
	userModelWriter
}

type userModelReader interface {
	GetByUsername(ctx context.Context, username string) (*User, error)
}

type userModelWriter interface {
	Insert(ctx context.Context, data *User) (int64, error)
}
