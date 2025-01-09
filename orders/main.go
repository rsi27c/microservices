package main

import (
	"example/orderservices/implementation"
	"example/orderservices/internals"
	"example/orderservices/pkg/loggers"
	"example/orderservices/pkg/oc"
	"example/orderservices/services"
	"example/orderservices/transports"
	net "example/orderservices/transports/http"

	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kithttp "github.com/go-kit/kit/transport/http"

	"log"
	"net/http"
	"os"
)

func main() {
	internals.GetEnv()
	loggers.OpenLog()
	db := internals.Connect()
	internals.Migrate(db)

	// server := fiber.New()

	repo := implementation.InitRepository(db)
	service := services.InitServices(repo)

	endpoint := transports.MakeEndpoints(service)

	endpoint = transports.Endpoints{
		CreateOrder:  oc.ServerEndpoint("CreateOrder")(endpoint.CreateOrder),
		GetOrder:     oc.ServerEndpoint("GetOrder")(endpoint.GetOrder),
		GetOrderByID: oc.ServerEndpoint("GetOrderByID")(endpoint.GetOrderByID),
		UpdateOrder:  oc.ServerEndpoint("UpdateOrder")(endpoint.UpdateOrder),
		DeleteOrder:  oc.ServerEndpoint("DeleteOrder")(endpoint.DeleteOrder),
	}
	ocTracing := kitoc.HTTPServerTrace()
	serverOptions := []kithttp.ServerOption{ocTracing}

	handler := net.NewService(endpoint, serverOptions)

	server := &http.Server{
		Addr:    os.Getenv("HTTP_PORT"),
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen to the server %v", err)
	}
}
