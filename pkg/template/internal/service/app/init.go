package app

import (
	"template/internal/app/hello"
	"template/internal/service/log"
)

func Init()  {
	hello.Init()

	log.Info("appService init successfully")
}
