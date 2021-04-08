package app

import (
	"conflict-template/internal/app/hello"
	"conflict-template/internal/service/log"
)

func Init() {
	hello.Init()

	log.Info("appService init successfully")
}
