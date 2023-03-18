package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type scoreSession struct {
	sqlc.CachedConn
	scoreRecord ScoreRecordSession
	userScore   UserScoreSession
}

func NewScoreSession(
	conn sqlx.SqlConn,
	c cache.CacheConf,
	scoreRecord ScoreRecordSession,
	userScore UserScoreSession,
) ScoreSession {
	return &scoreSession{
		CachedConn:  sqlc.NewConn(conn, c),
		scoreRecord: scoreRecord,
		userScore:   userScore,
	}
}

var _ ScoreSession = (*scoreSession)(nil)

func (m *scoreSession) UpdateUserScore(
	ctx context.Context,
	oldScore int32,
	userScore *UserScore,
	record *ScoreRecord,
) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := m.scoreRecord.CreateRecordSession(ctx, session,
			record); err != nil {
			return err
		}

		if err := m.userScore.UpdateUserScoreSession(ctx, session,
			oldScore, userScore); err != nil {
			return err
		}

		return nil
	})
}
