package graph

import "github.com/willychavez/order-listing-challenge/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	OrderUseCase usecase.OrderUseCase
}
