package web

import (
	"encoding/json"
	"net/http"

	"github.com/tiago-g-sales/clean-arch/internal/entity"
	"github.com/tiago-g-sales/clean-arch/internal/events"
	"github.com/tiago-g-sales/clean-arch/internal/usecase"
)




type WebOrderHandler struct{
	EventDispatcher events.EventDispatcherInterface
	OrderRepository entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}


func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreateEvent events.EventInterface,
) *WebOrderHandler{
	return &WebOrderHandler{
		EventDispatcher: EventDispatcher,
		OrderRepository: OrderRepository,
		OrderCreatedEvent: OrderCreateEvent,

	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request){
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}