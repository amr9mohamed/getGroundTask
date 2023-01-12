package tables

type Repository interface {
	Create(request CreateRequest) (Table, error)
}
