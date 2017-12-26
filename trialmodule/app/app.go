package app

import (
	"github.com/revel/revel"
	"github.com/revel/revel/logger"
)

var moduleLogger logger.MultiLogger

func init() {
	revel.OnAppStart(func() {
		moduleLogger.Info("Assigned Logger")
	})
	revel.RegisterModuleInit(func(module *revel.Module) {
		moduleLogger = module.Log
		moduleLogger.Info("Assigned Logger")
	})
}
