package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tasrifin/bookingtogo-golang/services"
	"github.com/gorilla/mux"
)

type CustomerController struct {
	customerService services.CustomerService
}

func NewCustomerController(service *services.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: *service,
	}
}

func (c *CustomerController) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	result := c.customerService.GetAllCustomers()

	respondWithJSON(w, result.Status, result.Payload)
}

func (c *CustomerController) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		error := map[string]interface{}{
			"error": "id must be integer - " + err.Error(),
		}
		respondWithJSON(w, http.StatusBadRequest, error)
		return
	}

	result := c.customerService.GetCustomerByID(id)

	respondWithJSON(w, result.Status, result.Payload)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
