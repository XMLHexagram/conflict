package service

import (
	"conflict-template/internal/service/app"
	"conflict-template/internal/service/config"
	"conflict-template/internal/service/db"
	"conflict-template/internal/service/httpEngine"
	"conflict-template/internal/service/log"
	"conflict-template/internal/service/toolkit"
)

func Init() {
	config.Init()
	log.Init()
	toolkit.Init()
	db.Init()
	httpEngine.Init()
	app.Init()
}

func Run() {
	httpEngine.Run()
}
