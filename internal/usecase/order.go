package usecase

import "github.com/willychavez/order-listing-challenge/internal/entity"

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type OrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *OrderUseCase {
	return &OrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *OrderUseCase) Execute(input OrderInputDTO) (OrderOutputDTO, error) {
	order := entity.Order{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}
	order.CalculateFinalPrice()
	if err := c.OrderRepository.Save(&order); err != nil {
		return OrderOutputDTO{}, err
	}

	dto := OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}

	return dto, nil
}

func (c *OrderUseCase) List() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetOrders()
	if err != nil {
		return nil, err
	}

	var dtos []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		dtos = append(dtos, dto)
	}

	return dtos, nil
}
