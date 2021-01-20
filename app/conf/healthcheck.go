package conf

import "time"

type Healthcheck struct {
	TTL  time.Duration `env:"HEALTHCHECK_TTL" envDefault:"1m"`
	MyIP string        `env:"MY_IP" envRequired:"true"`
}
