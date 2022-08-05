package app

import (
	"context"
	"embed"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"films-api/internal/api/delivery"
	"films-api/internal/config"
	"films-api/pkg/http/responder"
	"films-api/pkg/log"
)

type (
	// Meta defines meta for application.
	Meta struct {
		Info       Info
		ConfigPath string
	}

	// Info defines metadata of application.
	Info struct {
		AppName       string
		Tag           string
		Version       string
		Commit        string
		Date          string
		FortuneCookie string
	}

	// App defines main application struct.
	App struct {
		// meta information about application.
		meta Meta

		// tech dependencies.
		config *config.Config
		logger log.Logger

		dbMigrationsFS embed.FS
		db             *sqlx.DB
		rdb            *redis.Client

		responder responder.Responder

		// Delivery dependencies.
		statusHTTPHandler delivery.StatusHTTP
	}

	worker func(ctx context.Context, a *App)
)

// New - app constructor without init for components.
func New(meta Meta) *App {
	return &App{
		meta: meta,
	}
}

// WithMigrationFS is a setup for database migration filesystem
func (a *App) WithMigrationFS(f embed.FS) *App {
	a.dbMigrationsFS = f
	return a
}

// Run – registers graceful shutdown.
// populate configuration and application dependencies,
// run workers.
func (a *App) Run(ctx context.Context) {
	// Initialize configuration
	a.populateConfiguration()

	// Register Dependencies
	a.initLogger()
	a.initDatabase(ctx)

	// Domain registration.
	a.registerRepositories()
	a.registerServices()

	// Register Handlers
	a.registerHTTPHandlers()

	// Run Workers
	a.runWorkers(ctx)
}
