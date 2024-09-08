package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"github.com/tiago-g-sales/clean-arch/configs"
	"github.com/tiago-g-sales/clean-arch/internal/event/handler"
	"github.com/tiago-g-sales/clean-arch/internal/infra/web/webserver"
	"github.com/tiago-g-sales/clean-arch/pkg/events"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver , fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	createOrderUseCase.OrderCreated.GetName()

}


func getRabbitMQChannel() *amqp.Channel{

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	conn, err := amqp.Dial(fmt.Sprintf("%s://%s:%s@%s:%s/", configs.MQDriver, configs.MQUser, configs.MQHost, configs.MQPort))
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch

}

