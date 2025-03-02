package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"desafio_03/internal/entity"
	"desafio_03/internal/usecase"
	"desafio_03/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
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

func (h *WebOrderHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get("page")
	l := r.URL.Query().Get("limit")
	o := r.URL.Query().Get("order")

	var page uint = 0
	if p != "" {
		r, err := strconv.ParseUint(p, 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		page = uint(r)
	}

	var limit uint = 10
	if l != "" {
		r, err := strconv.ParseUint(l, 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		limit = uint(r)
	}

	if o != "" && o != "asc" && o != "desc" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	listOrder := usecase.NewListOrderUseCase(h.OrderRepository)

	paginationDto := usecase.PaginationDTO{Page: page, Limit: limit, Order: o}
	dto, err := listOrder.Execute(paginationDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
