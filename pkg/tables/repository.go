package tables

import (
	"github.com/getground/tech-tasks/backend/definitions/tables"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{db: db}
}

func (r repository) Create(req tables.CreateRequest) (t tables.Table, err error) {
	t = tables.Table{
		Capacity:   req.Capacity,
		EmptySeats: req.Capacity,
	}
	err = r.db.Create(&t).Error
	return
}
