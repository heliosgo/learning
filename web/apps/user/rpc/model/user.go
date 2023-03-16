package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

type User struct {
	Id         int64     `db:"id"`
	Username   string    `db:"username"`
	Password   string    `db:"password"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

const (
	cacheUserIdPrefix       = "cache:user:id:"
	cacheUserUsernamePrefix = "cache:user:username:"
)

var (
	userFieldNames        = builder.RawFieldNames(&User{})
	userRows              = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet = strings.Join(
		stringx.Remove(
			userFieldNames,
			"`id`",
			"`create_time`",
			"`update_time`",
			"`create_at`",
			"`update_at`",
		),
		",",
	)
)

type userModel struct {
	sqlc.CachedConn
	table string
}

var _ UserModel = (*userModel)(nil)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &userModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (u *userModel) GetByUsername(ctx context.Context, username string) (*User, error) {
	key := u.getUserUsernameKey(username)
	var res User
	err := u.QueryRowIndexCtx(
		ctx,
		&res,
		key,
		u.formatPrimary,
		func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (interface{}, error) {
			query := fmt.Sprintf(
				"select %s from %s where `username` = ? limit 1",
				userRows,
				u.table,
			)
			if err := conn.QueryRowCtx(ctx, &res, query, username); err != nil {
				return nil, err
			}

			return res.Id, nil
		},
		u.queryPrimary,
	)

	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &res, nil
}

func (u *userModel) Insert(ctx context.Context, data *User) (int64, error) {
	userUsernameKey := u.getUserUsernameKey(data.Username)
	res, err := u.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", u.table, userRowsExpectAutoSet)

		return conn.ExecCtx(ctx, query, data.Username, data.Password)
	}, userUsernameKey)

	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()

	return id, nil
}

func (u *userModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (u *userModel) queryPrimary(
	ctx context.Context,
	conn sqlx.SqlConn,
	v, primary interface{},
) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, u.tableName())

	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (u *userModel) tableName() string {
	return u.table
}

func (u *userModel) getUserUsernameKey(username string) string {
	return fmt.Sprintf("%s%s", cacheUserUsernamePrefix, username)
}

func (u *userModel) getUserIdKey(id int64) string {
	return fmt.Sprintf("%s%d", cacheUserIdPrefix, id)
}
