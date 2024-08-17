package events


// type EventInterface struct{

// }

type EventDispatcherInterface interface{
	Dispatcher(ed EventInterface)
}

type EventInterface interface{
	SetPayload(interface{})
	GetPayload() error
}
