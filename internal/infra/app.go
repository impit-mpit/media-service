package infra

import (
	"neuro-most/media-service/config"
	"neuro-most/media-service/internal/adapters/repo"
	"neuro-most/media-service/internal/infra/database"
	"neuro-most/media-service/internal/infra/router"
)

type app struct {
	cfg    config.Config
	router router.Router
	db     repo.GSQL
}

func Config(cfg config.Config) *app {
	return &app{cfg: cfg}
}

func (a *app) Database() *app {
	a.db = database.NewGormDB(a.cfg)
	return a
}

func (a *app) Serve() *app {
	a.router = router.NewRouter(a.db)
	return a
}

func (a *app) Start() {
	a.router.Listen()
}
