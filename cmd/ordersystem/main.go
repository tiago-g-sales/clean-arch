package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/streadway/amqp"
	"github.com/tiago-g-sales/clean-arch/configs"
	"github.com/tiago-g-sales/clean-arch/internal/event"
	"github.com/tiago-g-sales/clean-arch/internal/event/handler"
	"github.com/tiago-g-sales/clean-arch/internal/infra/database"
	"github.com/tiago-g-sales/clean-arch/internal/infra/graph"
	"github.com/tiago-g-sales/clean-arch/internal/infra/grpc/pb"
	"github.com/tiago-g-sales/clean-arch/internal/infra/grpc/service"
	"github.com/tiago-g-sales/clean-arch/internal/infra/web"
	"github.com/tiago-g-sales/clean-arch/internal/infra/web/webserver"
	"github.com/tiago-g-sales/clean-arch/internal/usecase"
	"github.com/tiago-g-sales/clean-arch/pkg/events"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	orderRepository := database.NewOrderRepository(db)
	err = orderRepository.CreateTableOrders()
	if err != nil {
		panic(err)
	}

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	orderCreated := event.NewOrderCreated()
	webOrderHandler := web.NewWebOrderHandler(eventDispatcher, orderRepository, orderCreated)
	createOrderUseCase := usecase.NewCreateOrderUseCase(orderRepository, orderCreated, eventDispatcher)
	listAllOrdersUseCase := usecase.NewListedOrderUseCase(orderRepository, orderCreated, eventDispatcher)

	

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webserver.AddHandler("POST", "/order", webOrderHandler.Create)
	webserver.AddHandler("GET", "/order",webOrderHandler.ListAll)



	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *listAllOrdersUseCase )

	pb.RegisterOrderServiceServer(grpcServer, createOrderService)

	
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	conn, err := amqp.Dial(fmt.Sprintf("%s://%s:%s@%s:%s/",configs.MQDriver ,configs.MQUser, configs.MQPassword, configs.MQHost, configs.MQPort ))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
