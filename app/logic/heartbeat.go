package logic

import (
	"github.com/petuhovskiy/yc-healthcheck/app/conf"
	"github.com/petuhovskiy/yc-healthcheck/app/store"
	"github.com/sirupsen/logrus"
	"time"
)

func StartHeartbeatLoop(repo *store.HealthcheckRepo, config conf.Healthcheck) {
	ticker := time.NewTicker(config.TTL / 2)

	for {
		h := store.Healthcheck{
			IP:     config.MyIP,
			Status: store.StatusAvailable,
			Moment: time.Now(),
		}

		err := repo.Save(h)
		if err != nil {
			logrus.WithError(err).Error("failed to save healthcheck")
		}

		<-ticker.C
	}
}
