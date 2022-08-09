package routes 

import (
	"github.com/gorilla/mux"
	"github.com/lyfeytech/MYSQL-ATCBank/pkg/controllers"
)

var RegisterBankRoutes = func (router *mux.Router){
	router.HandleFunc("/account/", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/account/", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/account/{customerId}", controllers.GetAccountById).Methods("GET")
	router.HandleFunc("/account/{customerId}", controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/account/{customerId}", controllers.DeleteAccount).Methods("DELETE")
	

}
