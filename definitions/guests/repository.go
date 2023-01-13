package guests

type Repository interface {
	Create(request CreateRequest, tableCapacity int64) error
}
