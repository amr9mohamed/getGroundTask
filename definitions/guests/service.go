package guests

type Service interface {
	Create(request CreateRequest) (CreateResponse, error)
}
