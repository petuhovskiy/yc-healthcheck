package logic

import "github.com/petuhovskiy/yc-healthcheck/app/store"

type Healthcheck struct {
	IP       string          `json:"ip"`
	Services []ServiceStatus `json:"services"`
}

type ServiceStatus struct {
	IP     string       `json:"ip"`
	Status store.Status `json:"status"`
}
