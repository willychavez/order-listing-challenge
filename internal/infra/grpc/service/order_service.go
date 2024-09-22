package service

import (
	"context"

	"github.com/willychavez/order-listing-challenge/internal/infra/grpc/pb"
	"github.com/willychavez/order-listing-challenge/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedListOrdersServer
	OrderUseCase usecase.OrderUseCase
}

func NewOrderService(orderUseCase usecase.OrderUseCase) *OrderService {
	return &OrderService{
		OrderUseCase: orderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.OrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.ListOrdersResponse, error) {
	output, err := s.OrderUseCase.List()
	if err != nil {
		return nil, err
	}
	var orders []*pb.OrderResponse
	for _, o := range output {
		orders = append(orders, &pb.OrderResponse{
			Id:         o.ID,
			Price:      float32(o.Price),
			Tax:        float32(o.Tax),
			FinalPrice: float32(o.FinalPrice),
		})
	}
	return &pb.ListOrdersResponse{
		Orders: orders,
	}, nil
}
