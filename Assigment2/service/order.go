package service

import (
	"Assigment2/models"
	"Assigment2/params"
	"Assigment2/repositories"
	"net/http"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
}

func NewOrderService(repo repositories.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: repo,
	}
}
func (o *OrderService) CreateOrder(request params.CreateOrder) (*params.Response, uint, error) {
	model := models.Order{
		CustomerName: request.CustomerName,
		Quantity:     request.Quantity,
	}
	orderId, err := o.orderRepo.CreateOrder(&model)
	if err != nil {
		return &params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REUEST",
			Error:   err.Error(),
		}, orderId, err
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: request,
	}, orderId, err
}
func (o *OrderService) GetOrders() (*params.Response, *[]models.Order, error) {
	orders, err := o.orderRepo.GetOrders()
	if err != nil {
		return &params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   err.Error(),
		}, nil, err
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: orders,
	}, orders, err
}
func (o *OrderService) DeleteOrders(orderId uint) (*params.Response, error) {
	err := o.orderRepo.DeleteOrders(orderId)
	if err != nil {
		return &params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   err.Error(),
		}, err
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: params.ResponseDeleteOrder{
			orderId: orderId,
		},
	}, err
}
func (o *OrderService) UpdateOrders(request params.UpdateOrder) (*params.Response, error) {
	var updateOrderReq = models.Order{
		ID:           request.ID,
		CustomerName: request.CustomerName,
	}
	err := o.orderRepo.UpdateOrders(&updateOrderReq)
	if err != nil {
		return &params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   err.Error(),
		}, err
	}
	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: params.UpdateOrder{
			CustomerName: request.CustomerName,
		},
	}, err
}
