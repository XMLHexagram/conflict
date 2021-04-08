package httpEngine

import (
	"github.com/gin-gonic/gin"
	"template/internal/service/config"
	"template/internal/service/log"
)

type service struct {
	HttpEngine *gin.Engine `wire:"-"`
	config.Http
}

var httpEngineService *service

func Init() {
	s := InitDep()

	s.HttpEngine = gin.Default()

	httpEngineService = s

	log.Info("httpEngineService init successfully")
}
