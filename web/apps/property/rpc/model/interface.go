package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type UserScoreModel interface {
	userScoreModelReader
	userScoreModelWriter

	UserScoreSession
}

type userScoreModelReader interface {
	GetByUserID(ctx context.Context, id int64) (*UserScore, error)
	InsertNew(ctx context.Context, uid int64) error
}

type userScoreModelWriter interface{}

type UserScoreSession interface {
	UpdateUserScoreSession(
		ctx context.Context,
		session sqlx.Session,
		oldScore int32,
		data *UserScore,
	) error
}

type LevelScoreModel interface {
	levelScoreModelReader
	levelScoreModelWriter
}

type levelScoreModelReader interface {
	FindAll(ctx context.Context) ([]*LevelScore, error)
}

type levelScoreModelWriter interface{}

type ScoreRecordModel interface {
	scoreRecordModelReader
	scoreRecordModelWriter

	ScoreRecordSession
}

type scoreRecordModelReader interface{}

type scoreRecordModelWriter interface{}

type ScoreRecordSession interface {
	CreateRecordSession(
		ctx context.Context,
		session sqlx.Session,
		data *ScoreRecord,
	) error
}

type ScoreSession interface {
	UpdateUserScore(
		ctx context.Context,
		oldScore int32,
		userScore *UserScore,
		record *ScoreRecord,
	) error
}
