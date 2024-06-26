package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/CodeBangkok/service"
	"github.com/gorilla/mux"
)

type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSrv.GetCustomers()
	if err != nil {
		handlerError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	customer, err := h.custSrv.GetCustomer(customerID)
	if err != nil {
		handlerError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer.CustomerID)
}
