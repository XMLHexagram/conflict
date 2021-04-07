package service

import (
	"template/internal/service/config"
	"template/internal/service/db"
	"template/internal/service/log"
)

func Init() {
	config.Init()
	log.Init()
	db.Init()
}
