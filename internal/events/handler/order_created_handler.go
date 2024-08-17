package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
	"github.com/tiago-g-sales/clean-arch/internal/events"
)

type OrderCreatedHandler struct{
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreateHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler{
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup){
	defer wg. Done()
	fmt.Printf("Order created: %v", event.GetPayload())
	jsonOutput, _:= json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body: jsonOutput,
	}
	h.RabbitMQChannel.Publish(
		"amq.direct", 
		"",
		false,
		false, 
		msgRabbitmq,
	)
}

