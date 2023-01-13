package guests

type Service interface {
	Create(request CreateRequest) (CreateResponse, error)
	GetGuestList() (ListDTO, error)
	GetGuests() (DTO, error)
	CheckIn(req CheckInRequest) (CheckInResponse, error)
	CheckOut(name string) error
}
