package tables

import (
	"errors"
	"github.com/getground/tech-tasks/backend/definitions/tables"
)

type Service struct {
	repository tables.Repository
}

func NewService(repository tables.Repository) Service {
	return Service{repository: repository}
}

func (s Service) Create(req tables.CreateRequest) (res tables.CreateResponse, err error) {
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

func (s Service) GetByID(id uint) (t tables.Table, err error) {
	t, err = s.repository.GetByID(id)
	if err != nil {
		err = errors.New("table not found")
	}
	return
}

func (s Service) CountEmptySeats() (count int) {
	return s.repository.CountEmptySeats()
}
