
package httpEngine

import (
	"github.com/gin-gonic/gin"
	"github.com/go-conflict/toolkit/hduhelp/sign"
	"github.com/go-conflict/toolkit/hduhelp/staff"
	"github.com/go-conflict/toolkit/httpHandler"
	"recruiting_system/internal/dep/config"
	"recruiting_system/internal/dep/log"
)

type service struct {
	HttpEngine *gin.Engine `wire:"-"`
	config.Http
}

var s *service

func init() {
	s = new(service)
	s.Http = config.ProvideHttp()

	s.HttpEngine = httpHandler.DefaultHttpEngine()
	s.HttpEngine.Use(sign.SignMiddleware(config.ProvideSign().Key),staff.ParseStaffInfoMiddleware())

	log.Info("httpEngineService init successfully")
}