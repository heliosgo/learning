package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserScore struct {
	UserID    int64     `db:"user_id"`
	Level     int32     `db:"level"`
	Score     int32     `db:"score"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type userScoreModel struct {
	sqlc.CachedConn
	table string
}

var _ UserScoreModel = (*userScoreModel)(nil)

func NewUserScoreModel(conn sqlx.SqlConn, c cache.CacheConf) UserScoreModel {
	return &userScoreModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_score`",
	}
}

func (m *userScoreModel) GetByUserID(ctx context.Context, id int64) (*UserScore, error) {
	query := fmt.Sprintf(
		"select user_id, level, score, created_at, updated_at from %s where user_id = ? limit 1",
		m.table,
	)
	var res UserScore
	if err := m.QueryRowNoCacheCtx(ctx, &res, query, id); err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &res, nil
}

func (m *userScoreModel) UpdateUserScoreSession(
	ctx context.Context,
	session sqlx.Session,
	oldScore int32,
	data *UserScore,
) error {
	query := fmt.Sprintf(
		"update %s set score = ?, level = ? where user_id = ? and score = ?",
		m.table,
	)

	res, err := session.ExecCtx(ctx, query, data.Score, data.Level, data.UserID, oldScore)
	if err != nil {
		return err
	}
	row, _ := res.RowsAffected()
	if row == 0 {
		return ErrNotFound
	}

	return nil
}

func (m *userScoreModel) InsertNew(ctx context.Context, uid int64) error {

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("insert into %s (user_id, level, score) values (?, ?, ?)", m.table)

		return conn.ExecCtx(ctx, query, uid, 0, 0)
	})

	return err
}
