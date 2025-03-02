package service

import (
	"context"
	"errors"

	"desafio_03/internal/infra/grpc/pb"
	"desafio_03/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	listOrderUseCase usecase.ListOrderUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrder(ctx context.Context, in *pb.PaginationRequest) (*pb.ListOrderResponse, error) {
	if in.Order != "" && in.Order != "asc" && in.Order != "desc" {
		return nil, errors.New("invalid order option")
	}

	if in.Limit == 0 {
		in.Limit = 10
	}

	paginationDto := usecase.PaginationDTO{
		Page:  uint(in.Page),
		Limit: uint(in.Limit),
		Order: in.Order,
	}
	dto, err := s.ListOrderUseCase.Execute(paginationDto)
	if err != nil {
		return nil, err
	}

	data := make([]*pb.OrderResponse, len(dto))

	for i, order := range dto {
		data[i] = &pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
	}

	return &pb.ListOrderResponse{Orders: data}, nil
}
