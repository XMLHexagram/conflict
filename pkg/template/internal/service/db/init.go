package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"template/internal/service/config"
	"template/internal/service/log"
)

var dbService *service

type service struct {
	DbMap       map[string]*gorm.DB
	DbConfigMap config.DbMap
}

func Init() {
	s := InitDep()

	for k, v := range s.DbConfigMap {
		dialector := mysql.Open(v.DSN)
		db, err := gorm.Open(dialector, &gorm.Config{})
		if err != nil {
			log.Panic("can not open db", zap.Error(err))
		}
		s.DbMap[k] = db
	}

	dbService = s

	log.Info("dbService init successfully")
}
