package events

import "github.com/tiago-g-sales/clean-arch/internal/usecase"

// type EventInterface struct{

// }

type EventDispatcherInterface interface{
	Dispatcher(ed EventInterface)
}

type EventInterface interface{
	SetPayload(dto usecase.OrderOutputDTO)
}
