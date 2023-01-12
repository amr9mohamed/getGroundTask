package tables

type Service interface {
	Create(request CreateRequest) (response CreateResponse, err error)
}
