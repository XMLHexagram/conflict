package service

import (
	"template/internal/service/app"
	"template/internal/service/config"
	"template/internal/service/db"
	"template/internal/service/httpEngine"
	"template/internal/service/log"
)

func Init() {
	config.Init()
	log.Init()
	db.Init()
	httpEngine.Init()
	app.Init()
}

func Run() {
	httpEngine.Run()
}
