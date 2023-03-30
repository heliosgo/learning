package model

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Task struct {
	Id        int32     `db:"id"`
	Name      string    `db:"name"`
	Desc      string    `db:"desc"`
	Type      int32     `db:"type"`
	Event     int32     `db:"event"`
	Target    int64     `db:"target"`
	Reward    string    `db:"reward"`
	Sort      int32     `db:"sort"`
	Status    int32     `db:"status"`
	JumpType  int32     `db:"jump_type"`
	JumpUri   string    `db:"jump_uri"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type tasksModel struct {
	sqlc.CachedConn
	table string
}

var _ TasksModel = (*tasksModel)(nil)

func NewTasksModel(conn sqlx.SqlConn, c cache.CacheConf) TasksModel {
	return &tasksModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tasks`",
	}
}

func (m *tasksModel) FindTasks(ctx context.Context) ([]*Task, error) {
	res := make([]*Task, 0)
	query := "select id, name, desc, type, event, target, reward, jump_type, jump_uri from tasks where status = ? order by sort"
	if err := m.QueryRowCtx(ctx, &res, "tasks", func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		return conn.QueryRowCtx(ctx, &res, query)
	}); err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return res, nil
		}
		return nil, err
	}

	return res, nil
}
