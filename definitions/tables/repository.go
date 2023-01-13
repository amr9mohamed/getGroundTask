package tables

type Repository interface {
	Create(request CreateRequest) (Table, error)
	GetByID(id uint) (Table, error)
	CountEmptySeats() int
}
