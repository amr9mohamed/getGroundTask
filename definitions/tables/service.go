package tables

type Service interface {
	Create(request CreateRequest) (response CreateResponse, err error)
	GetByID(id uint) (Table, error)
	CountEmptySeats() int
}
