package web

import (
	"encoding/json"
	"net/http"

	"github.com/willychavez/order-listing-challenge/internal/entity"
	"github.com/willychavez/order-listing-challenge/internal/usecase"
)

type WebOrderHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebOrderHandler(
	OrderRepository entity.OrderRepositoryInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		OrderRepository: OrderRepository,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewOrderUseCase(h.OrderRepository)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	getOrders := usecase.NewOrderUseCase(h.OrderRepository)
	output, err := getOrders.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) GetTotal(w http.ResponseWriter, r *http.Request) {
	total, err := h.OrderRepository.GetTotal()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(total)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
