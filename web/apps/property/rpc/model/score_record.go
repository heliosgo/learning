package model

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ScoreRecord struct {
	ID          int32     `db:"id"`
	UserID      int64     `db:"user_id"`
	ChangeScore int32     `db:"change_score"`
	AfterScore  int32     `db:"after_score"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type scoreRecordModel struct {
	sqlc.CachedConn
	table string
}

var _ ScoreRecordModel = (*scoreRecordModel)(nil)

func NewScoreRecordModel(conn sqlx.SqlConn, c cache.CacheConf) ScoreRecordModel {
	return &scoreRecordModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`score_record`",
	}
}

func (m *scoreRecordModel) CreateRecordSession(
	ctx context.Context,
	session sqlx.Session,
	data *ScoreRecord,
) error {
	query := fmt.Sprintf(
		"insert into %s (user_id, change_score, after_score) values (?, ?, ?)",
		m.table,
	)

	res, err := session.ExecCtx(ctx, query, data.UserID, data.ChangeScore, data.AfterScore)
	if err != nil {
		return err
	}
	row, _ := res.RowsAffected()
	if row == 0 {
		return ErrNotFound
	}

	return nil
}
