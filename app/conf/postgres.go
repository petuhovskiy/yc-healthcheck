package conf

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Postgres struct {
	DSN                string `env:"POSTGRES_DSN" envDefault:"host=localhost port=5432 sslmode=disable user=postgres dbname=postgres password="`
	MaxOpenConnections int    `env:"POSTGRES_MAX_OPEN_CONNECTIONS" envDefault:"20"`
	MaxIdleConnections int    `env:"POSTGRES_MAX_IDLE_CONNECTIONS" envDefault:"5"`
}

func (p Postgres) Connect() (*gorm.DB, error) {
	conn, err := gorm.Open("postgres", p.DSN)
	if err != nil {
		return nil, err
	}

	if logrus.GetLevel() >= logrus.DebugLevel {
		conn.SetLogger(logrus.StandardLogger())
		conn.LogMode(true)
	}

	conn.DB().SetMaxOpenConns(p.MaxOpenConnections)
	conn.DB().SetMaxIdleConns(p.MaxIdleConnections)

	return conn, nil
}
