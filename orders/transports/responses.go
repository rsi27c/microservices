package transports

import "example/orderservices/pkg/models"

type CreateRequest struct {
	Orders models.Order
}

type CreateResponse struct {
	Message string `json:"message,omitempty"`
	Error   error  `json:"error,omitempty"`
}

type GetOrderRequest struct{}

type GetOrderResponse struct {
	Orders []models.Order `json:"orders,omitempty"`
	Error  error          `json:"error,omitempty"`
}

type GetByIDRequest struct {
	ID int
}

type GetOrderByIDResponse struct {
	Orders models.Order `json:"orders,omitempty"`
	Error  error        `json:"error,omitempty"`
}

type UpdateOrderRequest struct {
	ID     int
	Orders models.Order
}

type UpdateOrderResponse struct {
	ID    int   `json:"id,omitempty"`
	Error error `json:"error,omitempty"`
}

type DeleteRequest struct {
	ID int
}

type DeleteOrderResponse struct {
	ID    int   `json:"id,omitempty"`
	Error error `json:"error,omitempty"`
}
