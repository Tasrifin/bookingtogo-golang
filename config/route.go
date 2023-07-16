package config

import (
	"net/http"

	"github.com/Tasrifin/bookingtogo-golang/controllers"
	"github.com/Tasrifin/bookingtogo-golang/repositories"
	"github.com/Tasrifin/bookingtogo-golang/services"
	"github.com/gorilla/mux"
)

func InitRoute() {
	db := ConnectDB()
	route := mux.NewRouter()

	//CUSTOMER REPO
	customerRepo := repositories.NewCustomerRepo(db)
	custtomerService := services.NewBattleService(customerRepo)
	customerContoller := controllers.NewCustomerController(custtomerService)

	route.HandleFunc("/customers", customerContoller.GetAllCustomers).Methods(http.MethodGet)
	route.HandleFunc("/customer/{id}", customerContoller.GetCustomerByID).Methods(http.MethodGet)

	http.ListenAndServe(":8080", route)
}
