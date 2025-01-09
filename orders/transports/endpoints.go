package transports

import (
	"context"
	"example/orderservices/services"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateOrder  endpoint.Endpoint
	GetOrder     endpoint.Endpoint
	GetOrderByID endpoint.Endpoint
	UpdateOrder  endpoint.Endpoint
	DeleteOrder  endpoint.Endpoint
}

func MakeEndpoints(s services.Services) Endpoints {
	return Endpoints{
		CreateOrder:  MakeCreateEndpoint(s),
		GetOrder:     MakeGetOrdersEndpoint(s),
		GetOrderByID: MakeGetOrderByIDEndpoint(s),
		UpdateOrder:  MakeupdateOrderEndpoint(s),
		DeleteOrder:  MakeDeleteOrderEndpoint(s),
	}
}

func MakeCreateEndpoint(service services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest) // type assertion
		err := service.CreateOrder(&req.Orders)
		if err != nil {
			return CreateResponse{Error: err}, err
		}
		return CreateResponse{Message: "Order placed successfully"}, nil
	}
}

func MakeGetOrdersEndpoint(service services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		orders, err := service.GetOrder()
		if err != nil {
			return GetOrderResponse{Error: err}, err
		}
		return GetOrderResponse{Orders: *orders}, nil
	}
}

func MakeGetOrderByIDEndpoint(service services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest) // type assertion
		orders, err := service.GetOrderByID(req.ID)
		if err != nil {
			return GetOrderByIDResponse{Error: err}, err
		}
		return GetOrderByIDResponse{Orders: *orders}, nil
	}
}

func MakeupdateOrderEndpoint(service services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateOrderRequest) // type assertion
		err := service.UpdateOrder(&req.Orders, req.ID)
		if err != nil {
			return UpdateOrderResponse{Error: err}, err
		}
		return UpdateOrderResponse{ID: req.ID}, nil
	}
}

func MakeDeleteOrderEndpoint(service services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest) // type assertion
		err := service.DeleteOrder(req.ID)
		if err != nil {
			return DeleteOrderResponse{Error: err}, err
		}
		return DeleteOrderResponse{ID: req.ID}, nil
	}
}
