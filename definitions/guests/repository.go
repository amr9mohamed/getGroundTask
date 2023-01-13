package guests

import "github.com/getground/tech-tasks/backend/definitions/tables"

type Repository interface {
	Create(request CreateRequest, tableCapacity int64) error
	GetByName(name string) (Guest, error)
	GetGuestList(arrived bool) ([]Guest, error)
	CheckIn(request CheckInRequest, guest Guest, table tables.Table) error
	CheckOut(name string) error
}
