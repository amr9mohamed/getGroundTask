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

func (r Repository) GetByName(name string) (g guests.Guest, err error) {
	err = r.db.Where("name = ?", name).Where("time_arrived IS NULL").First(&g).Error
	return
}

func (r Repository) GetGuestList(arrived bool) (list []guests.Guest, err error) {
	if arrived {
		err = r.db.Where("time_arrived IS NOT NULL").Find(&list).Error
		return
	}
	err = r.db.Find(&list).Error // get all arrived or not
	return
}

func (r Repository) CheckIn(req guests.CheckInRequest, g guests.Guest, t tables.Table) (err error) {
	return r.db.Transaction(
		func(tx *gorm.DB) error {
			ts := time.Now()
			err := tx.Where(&guests.Guest{Name: req.Name}).Updates(
				guests.Guest{
					TimeArrived:  &ts,
					Accompanying: req.Accompanying,
				},
			).Error
			if err != nil {
				return err
			}

			err = tx.
				Where(&tables.Table{ID: t.ID}).
				Select("capacity", "empty_seats").
				Updates(
					tables.Table{
						Capacity:   t.Capacity - (req.Accompanying - g.Accompanying),
						EmptySeats: t.EmptySeats - req.Accompanying - 1,
					},
				).
				Error
			return err
		},
	)
}

func (r Repository) CheckOut(name string) (err error) {
	// check if guest exists and already checked in
	g := guests.Guest{}
	err = r.db.
		Where("name = ?", name).
		Where("checked_out = 0").
		Where("time_arrived IS NOT NULL").
		First(&g).Error
	if err != nil {
		return err
	}
	t := tables.Table{ID: g.TableID}
	err = r.db.First(&t).Error
	if err != nil {
		return err
	}

	tx := r.db.Begin()
	// set guest checked out flag
	err = tx.Where(&guests.Guest{Name: name}).Updates(guests.Guest{CheckedOut: 1}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// update empty seats
	err = tx.Where(&tables.Table{ID: t.ID}).Updates(tables.Table{EmptySeats: t.EmptySeats + g.Accompanying + 1}).Error
	if err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}
