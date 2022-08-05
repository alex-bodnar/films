package app

import (
	"films-api/pkg/log"
)

// initLogger initializes logger.
func (a *App) initLogger() {
	a.logger = log.InitLogger(a.config.Logger, map[string]string{
		"service": a.meta.Info.AppName,
	})
}
