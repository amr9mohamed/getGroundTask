package tables

import "github.com/getground/tech-tasks/backend/definitions/tables"

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

func (s service) CountEmptySeats() (count int) {
	return s.repository.CountEmptySeats()
}
