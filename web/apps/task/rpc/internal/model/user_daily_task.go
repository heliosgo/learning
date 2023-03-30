package model

import (
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserDailyTask struct {
	Id           int32     `db:"id"`
	UserId       int64     `db:"user_id"`
	TaskId       int32     `db:"task_id"`
	TaskType     int32     `db:"task_type"`
	TaskEvent    int32     `db:"task_event"`
	TaskTarget   int64     `db:"task_target"`
	TaskProgress int64     `db:"task_progress"`
	Status       int32     `db:"status"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type userDailyTaskModel struct {
	sqlc.CachedConn
	tablePrefix string
}

var _ UserDailyTaskModel = (*userDailyTaskModel)(nil)

func NewUserDailyTaskModel(conn sqlx.SqlConn, c cache.CacheConf) UserDailyTaskModel {
	return &userDailyTaskModel{
		CachedConn:  sqlc.NewConn(conn, c),
		tablePrefix: "user_daily_task_",
	}
}

func (m *userDailyTaskModel) getTableName(t time.Time) string {
	return fmt.Sprintf("`%s%s`", m.tablePrefix, t.Format("20060102"))
}
