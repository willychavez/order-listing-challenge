package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/willychavez/order-listing-challenge/config"
	"github.com/willychavez/order-listing-challenge/container"
	"github.com/willychavez/order-listing-challenge/internal/infra/graph"
	"github.com/willychavez/order-listing-challenge/internal/infra/grpc/pb"
	"github.com/willychavez/order-listing-challenge/internal/infra/grpc/service"
	"github.com/willychavez/order-listing-challenge/internal/infra/web/webserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.Get()

	db, err := sql.Open(cfg.DB.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	webserver := webserver.NewWebServer(cfg.HTTP.Port)
	webOrderHandler := container.NewWebOrderHandler(db)
	webserver.AddHandler("POST /order", webOrderHandler.Create)
	webserver.AddHandler("GET /order", webOrderHandler.GetOrders)
	go webserver.Start()

	orderUseCase := container.NewOrderUseCase(db)

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*orderUseCase)
	pb.RegisterListOrdersServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", cfg.GRPC.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC.Port))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		OrderUseCase: *orderUseCase,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.GraphQL.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.GraphQL.Port, nil))

}
