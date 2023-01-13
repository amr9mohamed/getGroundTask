package guests

import (
	"errors"
	"github.com/getground/tech-tasks/backend/definitions/guests"
	"github.com/getground/tech-tasks/backend/definitions/tables"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Create(req guests.CreateRequest, tableCapacity int64) (err error) {
	return r.db.Transaction(
		func(tx *gorm.DB) error {
			// create guest
			err := r.db.Model(&guests.Guest{}).Create(
				guests.Guest{
					Name:         req.Name,
					TableID:      req.Table,
					Accompanying: req.Accompanying,
					TimeArrived:  time.Now(),
				},
			).Error
			if err != nil {
				return errors.New(err.Error())
			}

			// update table capacity if guest created
			err = r.db.Where(&tables.Table{ID: req.Table}).Updates(tables.Table{Capacity: tableCapacity}).Error
			if err != nil {
				return errors.New("error reserving table seats")
			}
			return nil
		},
	)
}
