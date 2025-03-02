package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	// GetTotal() (int, error)
	ListAll(offset uint, limit uint, order string) ([]Order, error)
}
