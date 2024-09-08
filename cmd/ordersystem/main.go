package main

import (
	"database/sql"
	"fmt"
	"net"

	_"github.com/99designs/gqlgen/graphql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"github.com/tiago-g-sales/clean-arch/configs"
	"github.com/tiago-g-sales/clean-arch/internal/event/handler"
	_"github.com/tiago-g-sales/clean-arch/internal/infra/graph"
	"github.com/tiago-g-sales/clean-arch/internal/infra/grpc/pb"
	"github.com/tiago-g-sales/clean-arch/internal/infra/grpc/service"
	"github.com/tiago-g-sales/clean-arch/internal/infra/web/webserver"
	"github.com/tiago-g-sales/clean-arch/pkg/events"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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


	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)
	
	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)



}


func getRabbitMQChannel() *amqp.Channel{

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	conn, err := amqp.Dial(fmt.Sprintf("%s://%s:%s@%s:%s/", configs.MQDriver, configs.MQUser, configs.MQUser, configs.MQHost, configs.MQPort))
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch

}

