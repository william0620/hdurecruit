package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"recruiting_system/internal/dep/config"
	"recruiting_system/internal/dep/log"
)

type service struct {
	DbMap       map[string]*gorm.DB
	DbConfigMap config.DbMap
}

var s *service

func init()  {
	s = new(service)
	s.DbMap = make(map[string]*gorm.DB,100)
	s.DbConfigMap = config.ProvideDbMap()

	for k, v := range s.DbConfigMap {
		dialector := mysql.Open(v.DSN)
		db, err := gorm.Open(dialector, &gorm.Config{})
		if err != nil {
			log.Panic("can not open db", zap.Error(err))
		}
		s.DbMap[k] = db
	}

	log.Info("dbService init successfully")
}