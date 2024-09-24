package usecase

import (
	"github.com/tiago-g-sales/clean-arch/internal/domain"
	"github.com/tiago-g-sales/clean-arch/internal/entity"
	"github.com/tiago-g-sales/clean-arch/pkg/events"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}


func NewListedOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderListed:    OrderListed,
		EventDispatcher: EventDispatcher,
	}
}


func (c *ListOrderUseCase) Execute() ( []*domain.OrderOutputDTO, error) {

	listOrders, err := c.OrderRepository.FindAll()
	if err != nil {
		return []*domain.OrderOutputDTO{}, err
	}
	var ordersResponse []*domain.OrderOutputDTO

	for _, order := range listOrders {
		dto :=domain.OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}	
		ordersResponse = append(ordersResponse, &dto)
	} 


	//c.OrderListed.SetPayload(ordersResponse)
	//c.EventDispatcher.Dispatch(c.OrderListed)

	return ordersResponse, nil
}
