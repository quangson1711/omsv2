package main

import (
	"common"
	pb "common/api"
	"net/http"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(writer http.ResponseWriter, request *http.Request) {
	customerID := request.PathValue("customerID")

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(request, &items); err != nil {
		common.WirteError(writer, http.StatusBadRequest, err.Error())
	}

	h.client.CreateOrder(request.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
}
