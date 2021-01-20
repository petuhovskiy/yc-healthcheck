package logic

import (
	"github.com/petuhovskiy/yc-healthcheck/app/store"
	"github.com/sirupsen/logrus"
	"time"
)

type Service struct {
	repo *store.HealthcheckRepo
	myIP string
	ttl  time.Duration
}

func NewService(repo *store.HealthcheckRepo, myIP string, ttl time.Duration) *Service {
	return &Service{
		repo: repo,
		myIP: myIP,
		ttl:  ttl,
	}
}

func (s *Service) computeStatus(h store.Healthcheck) store.Status {
	if time.Since(h.Moment) >= s.ttl {
		return store.StatusNotAvailable
	}

	return h.Status
}

func (s *Service) Healthcheck() (*Healthcheck, error) {
	all, err := s.repo.FindAll()
	if err != nil {
		logrus.WithError(err).Error("failed to find all healthchecks")
		return nil, ErrDatabaseNotAvailable
	}

	var statuses []ServiceStatus
	for _, h := range all {
		statuses = append(
			statuses,
			ServiceStatus{
				IP:     h.IP,
				Status: s.computeStatus(h),
			},
		)
	}

	return &Healthcheck{
		IP:       s.myIP,
		Services: statuses,
	}, nil
}
