package dal

type OrderRepository struct {
	filePath string
}

func NewOrderRepository(filePath string) *OrderRepository {
	return &OrderRepository{filePath: filePath}
}
