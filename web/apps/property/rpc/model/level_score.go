package model

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type LevelScore struct {
	Level int32 `db:"level"`
	Score int32 `db:"score"`
}

type levelScoreModel struct {
	sqlc.CachedConn
	table string
}

var _ LevelScoreModel = (*levelScoreModel)(nil)

func NewLevelScoreModel(conn sqlx.SqlConn, c cache.CacheConf) LevelScoreModel {
	return &levelScoreModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`level_score`",
	}
}

func (m *levelScoreModel) FindAll(ctx context.Context) ([]*LevelScore, error) {
	res := make([]*LevelScore, 0)
	query := "select level, score from level_score order by level asc"
	if err := m.QueryRowsNoCacheCtx(ctx, &res, query); err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return res, ErrNotFound
		}

		return res, err
	}

	return res, nil
}
