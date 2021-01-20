package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/petuhovskiy/yc-healthcheck/app/conf"
	"github.com/petuhovskiy/yc-healthcheck/app/logic"
	"github.com/petuhovskiy/yc-healthcheck/app/server"
	"github.com/petuhovskiy/yc-healthcheck/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	cfg, err := conf.ParseEnv()
	if err != nil {
		logrus.WithError(err).Fatal("failed to parse config")
	}

	db, err := cfg.Postgres.Connect()
	if err != nil {
		logrus.WithError(err).Fatal("failed to connect to db")
	}

	db.AutoMigrate(store.Healthcheck{})

	healthcheckRepo := store.NewHealthcheckRepo(db)

	// start background healthcheck loop
	go logic.StartHeartbeatLoop(healthcheckRepo, cfg.Healthcheck)

	serv := logic.NewService(healthcheckRepo, cfg.Healthcheck.MyIP, cfg.Healthcheck.TTL)
	handler := server.NewHandler(serv)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/healthcheck", handler.Healthcheck)

	err = http.ListenAndServe(cfg.Server.Bind, r)
	if err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Fatal("main server crashed")
	}

	logrus.Info("main server finished")
}
