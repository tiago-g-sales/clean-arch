package service

import (
	"context"

	"github.com/tiago-g-sales/clean-arch/internal/domain"
	"github.com/tiago-g-sales/clean-arch/internal/infra/grpc/pb"
	"github.com/tiago-g-sales/clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  usecase.CreateOrderUseCase
	ListAllOrderUseCase usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listAllOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		ListAllOrderUseCase: listAllOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	dto := domain.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
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

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.ListAllOrdersResponse, error) {

	listallorders, err := s.ListAllOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.OrderResponse

	for _, order := range listallorders {
		dto := pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
		ordersResponse = append(ordersResponse, &dto)
	}

	return &pb.ListAllOrdersResponse{
		Orderresponse: ordersResponse}, nil
}
