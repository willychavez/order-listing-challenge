package container

import (
	"database/sql"

	"github.com/willychavez/order-listing-challenge/internal/infra/database"
	"github.com/willychavez/order-listing-challenge/internal/infra/web"
	"github.com/willychavez/order-listing-challenge/internal/usecase"
)

// NewOrdersUseCase creates an instance of GetOrdersUseCase with the necessary dependencies
func NewOrderUseCase(db *sql.DB) *usecase.OrderUseCase {
	orderRepository := database.NewOrderRepository(db)
	return usecase.NewOrderUseCase(orderRepository)
}

// NewWebOrderHandler creates an instance of WebOrderHandler with the necessary dependencies
func NewWebOrderHandler(db *sql.DB) *web.WebOrderHandler {
	orderRepository := database.NewOrderRepository(db)
	return web.NewWebOrderHandler(orderRepository)
}
