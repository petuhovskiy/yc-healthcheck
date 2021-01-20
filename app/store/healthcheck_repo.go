package store

import "github.com/jinzhu/gorm"

type HealthcheckRepo struct {
	db *gorm.DB
}

func NewHealthcheckRepo(db *gorm.DB) *HealthcheckRepo {
	return &HealthcheckRepo{
		db: db,
	}
}

func (r *HealthcheckRepo) Save(h Healthcheck) error {
	return r.db.Save(&h).Error
}

func (r *HealthcheckRepo) FindAll() ([]Healthcheck, error) {
	var res []Healthcheck
	err := r.db.Find(&res).Error
	return res, err
}
