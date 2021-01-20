package store

import "time"

type Healthcheck struct {
	IP     string `gorm:"primary_key"`
	Status Status
	Moment time.Time
}
