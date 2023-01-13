package guests

import (
	"errors"
	"github.com/getground/tech-tasks/backend/definitions/guests"
	"github.com/getground/tech-tasks/backend/definitions/tables"
)

type Service struct {
	repository guests.Repository
	tableSvc   tables.Service
}

func NewService(repository guests.Repository, tableSvc tables.Service) Service {
	return Service{repository: repository, tableSvc: tableSvc}
}

func (s Service) Create(req guests.CreateRequest) (res guests.CreateResponse, err error) {
	t, err := s.tableSvc.GetByID(req.Table)
	if err != nil {
		return
	}

	// validate capacity, can be moved to validator
	newTableCapacity := t.Capacity - req.Accompanying - 1
	if newTableCapacity < 0 {
		err = errors.New("table have no capacity for accompanying")
		return
	}

	err = s.repository.Create(req, newTableCapacity)
	if err != nil {
		return
	}
	res.Name = req.Name
	return
}

func (s Service) GetGuestList() (list guests.ListDTO, err error) {
	res, err := s.repository.GetGuestList(false)
	if err != nil {
		return
	}
	list = mapGuestsListToDTO(res)
	return
}

func (s Service) GetGuests() (list guests.DTO, err error) {
	res, err := s.repository.GetGuestList(true)
	if err != nil {
		return
	}
	list = mapGuestsToDTO(res)
	return
}
