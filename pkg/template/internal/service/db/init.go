package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"template/internal/service/config"
	"template/internal/service/log"
)

var dbService *service

type service struct {
	DefaultDb *gorm.DB `wire:"-"`
	DbMap     config.DbMap
}

func Init() {
	s := InitDep()

	dsnDefault := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		s.DbMap["default"].User,
		s.DbMap["default"].Password,
		s.DbMap["default"].Address,
		s.DbMap["default"].DBName,
	)

	NekoMySQL, err := gorm.Open(mysql.Open(dsnDefault), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	s.DefaultDb = NekoMySQL

	dbService = s
	log.Info("dbService init successfully")
}
