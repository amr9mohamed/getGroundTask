package guests

type Service interface {
	Create(request CreateRequest) (CreateResponse, error)
	GetGuestList() (ListDTO, error)
	GetGuests() (DTO, error)
}
