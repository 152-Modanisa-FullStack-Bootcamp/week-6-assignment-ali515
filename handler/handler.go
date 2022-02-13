package handler

import (
	"bootcamp/model"
	"bootcamp/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type IHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.IService
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	username := strings.Split(r.URL.Path, "/")[1]
	customer, err := h.service.GetCustomerByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		return
	}
	result, err := json.Marshal(customer)
	w.Write(result)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	customers := h.service.GetAllCustomers()
	if customers == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	customerResponse := model.CustomerResponse{}
	for _, v := range *customers {
		customerResponse = append(customerResponse, v)
	}
	result, err := json.Marshal(customerResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if customer.Username == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	err = h.service.CreateCustomer(&customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, _ := json.Marshal(true)
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if customer.Username == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	oldCustomer, err := h.service.GetCustomerByUsername(customer.Username)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		return
	}
	oldCustomer.Name = customer.Name
	oldCustomer.Surname = customer.Surname
	err = h.service.UpdateCustomer(oldCustomer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, _ := json.Marshal(true)
	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}

func NewHandler(service service.IService) IHandler {
	return &handler{service: service}
}
