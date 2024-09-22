package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"

	"github.com/willychavez/order-listing-challenge/internal/infra/graph/model"
	"github.com/willychavez/order-listing-challenge/internal/usecase"
)

// CreateOrder is the resolver for the CreateOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}
	output, err := r.OrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// ListOrders is the resolver for the ListOrders field.
func (r *queryResolver) ListOrders(ctx context.Context) ([]*model.Order, error) {
	output, err := r.OrderUseCase.List()
	if err != nil {
		return nil, err
	}

	var orders []*model.Order
	for _, o := range output {
		orders = append(orders, &model.Order{
			ID:         o.ID,
			Price:      float64(o.Price),
			Tax:        float64(o.Tax),
			FinalPrice: float64(o.FinalPrice),
		})
	}
	return orders, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
