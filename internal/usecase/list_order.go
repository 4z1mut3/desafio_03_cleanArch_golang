package usecase

import (
	"desafio_03/internal/entity"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute(p PaginationDTO) ([]OrderOutputDTO, error) {
	var dto []OrderOutputDTO

	orders, err := c.OrderRepository.ListAll(p.Page, p.Limit, p.Order)
	if err != nil {
		return nil, err
	}

	for _, o := range orders {
		dto = append(dto, OrderOutputDTO{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}

	return dto, nil
}
