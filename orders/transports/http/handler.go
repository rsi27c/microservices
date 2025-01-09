package http

import (
	"context"
	"encoding/json"
	"example/orderservices/services"
	"example/orderservices/transports"
	"fmt"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	services.Services
}

func NewService(svcEndpoints transports.Endpoints, options []kithttp.ServerOption) http.Handler {
	// set-up router and initialize http endpoints
	var (
		r            = mux.NewRouter()
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)
	options = append(options, errorEncoder)
	//options := []kithttp.ServerOption{
	//	kithttp.ServerErrorLogger(logger),
	//	kithttp.ServerErrorEncoder(encodeError),
	//}
	// HTTP Post - /orders
	r.Methods("POST").Path("/orders").Handler(kithttp.NewServer(
		svcEndpoints.CreateOrder,
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	// HTTP Post - /orders
	r.Methods("GET").Path("/orders").Handler(kithttp.NewServer(
		svcEndpoints.GetOrder,
		decodeGetOrdersRequest,
		encodeResponse,
		options...,
	))

	// HTTP Post - /orders/{id}
	r.Methods("GET").Path("/orders/{id}").Handler(kithttp.NewServer(
		svcEndpoints.GetOrderByID,
		decodeGetByIDRequest,
		encodeResponse,
		options...,
	))

	// HTTP Post - /orders/status
	r.Methods("PUT").Path("/orders/{id}").Handler(kithttp.NewServer(
		svcEndpoints.UpdateOrder,
		decodeUpdateRequest,
		encodeResponse,
		options...,
	))

	// HTTP Post - /orders/{id}
	r.Methods("DELETE").Path("/orders/{id}").Handler(kithttp.NewServer(
		svcEndpoints.DeleteOrder,
		decodeDeleteRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transports.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Orders); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetOrdersRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return transports.GetOrderRequest{}, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	idstr, ok := vars["id"]
	if !ok {
		return nil, fmt.Errorf("bad routing")
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return nil, err
	}

	return transports.GetByIDRequest{ID: id}, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	idstr, ok := vars["id"]
	if !ok {
		return nil, fmt.Errorf("bad routing")
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return nil, err
	}

	var req transports.UpdateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return transports.UpdateOrderRequest{ID: id, Orders: req.Orders}, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	idstr, ok := vars["id"]
	if !ok {
		return nil, fmt.Errorf("bad routing")
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return nil, err
	}

	return transports.DeleteRequest{ID: id}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case gorm.ErrEmptySlice:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
