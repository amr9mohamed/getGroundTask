package tables

import (
	"errors"
	"github.com/getground/tech-tasks/backend/definitions/tables"
)

type service struct {
	repository tables.Repository
}

func NewService(repository tables.Repository) service {
	return service{repository: repository}
}

func (s service) Create(req tables.CreateRequest) (res tables.CreateResponse, err error) {
	t, err := s.repository.Create(req)
	if err != nil {
		return
	}
	res = tables.CreateResponse{
		ID:       t.ID,
		Capacity: t.Capacity,
	}
	return
}

func (s service) GetByID(id uint) (t tables.Table, err error) {
	t, err = s.repository.GetByID(id)
	if err != nil {
		err = errors.New("table not found")
	}
	return
}

func (s service) CountEmptySeats() (count int) {
	return s.repository.CountEmptySeats()
}
