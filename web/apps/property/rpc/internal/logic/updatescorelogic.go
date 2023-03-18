package logic

import (
	"context"

	"web/apps/property/rpc/internal/svc"
	"web/apps/property/rpc/model"
	"web/apps/property/rpc/property"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScoreLogic {
	return &UpdateScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateScoreLogic) UpdateScore(
	in *property.UpdateScoreReq,
) (*property.UpdateScoreRsp, error) {
	userScore, err := l.svcCtx.UserScoreModel.GetByUserID(l.ctx, in.Uid)
	if err != nil {
		return nil, errors.Wrapf(err, "get user score failed, uid: %d", in.Uid)
	}

	newScore := userScore.Score + in.Score

	i, j := 0, len(l.svcCtx.LevelScoreSli)-1
	for i < j {
		m := (i + j + 1) >> 1
		if l.svcCtx.LevelScoreSli[m].Score <= newScore {
			i = m
		} else {
			j = m - 1
		}
	}
	userScore.Level = l.svcCtx.LevelScoreSli[i].Level

	if err := l.svcCtx.ScoreSession.UpdateUserScore(
		l.ctx,
		userScore.Score,
		&model.UserScore{
			UserID: userScore.UserID,
			Score:  newScore,
			Level:  userScore.Level,
		}, &model.ScoreRecord{
			UserID:      userScore.UserID,
			ChangeScore: in.Score,
			AfterScore:  newScore,
		}); err != nil {

		return nil, errors.Wrapf(
			err,
			"update user score failed, uid: %d, delta score: %d",
			userScore.UserID,
			in.Score,
		)
	}

	return &property.UpdateScoreRsp{
		Score: newScore,
		Level: userScore.Level,
	}, nil
}
