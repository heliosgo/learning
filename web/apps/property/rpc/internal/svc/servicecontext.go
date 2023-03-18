package svc

import (
	"context"
	"log"
	"time"
	"web/apps/property/rpc/internal/config"
	"web/apps/property/rpc/model"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	LocalCache    *collection.Cache
	LevelScoreSli []*model.LevelScore

	LevelScoreModel  model.LevelScoreModel
	UserScoreModel   model.UserScoreModel
	ScoreRecordModel model.ScoreRecordModel

	ScoreSession model.ScoreSession
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	localCache, err := collection.NewCache(time.Minute, collection.WithLimit(10))
	if err != nil {
		log.Fatal(err)
	}
	levelScoreModel := model.NewLevelScoreModel(sqlConn, c.CacheRedis)
	sli, err := levelScoreModel.FindAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	userScoreModel := model.NewUserScoreModel(sqlConn, c.CacheRedis)
	scoreRecordModel := model.NewScoreRecordModel(sqlConn, c.CacheRedis)
	return &ServiceContext{
		Config:           c,
		LocalCache:       localCache,
		LevelScoreSli:    sli,
		LevelScoreModel:  levelScoreModel,
		UserScoreModel:   userScoreModel,
		ScoreRecordModel: scoreRecordModel,
		ScoreSession: model.NewScoreSession(
			sqlConn,
			c.CacheRedis,
			scoreRecordModel,
			userScoreModel,
		),
	}
}
