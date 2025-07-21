package routes

import (
	"github.com/fakh1m/go-auth-invoice/controller"
	"github.com/fakh1m/go-auth-invoice/middleware"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("api/register", controller.Register).Methods("POST")
	r.HandleFunc("/api/login", controller.Login).Methods("POST")

	invoiceRouter := r.PathPrefix("/api/invoices").Subrouter()
	invoiceRouter.Use(middleware.JWTMiddleware)
	invoiceRouter.HandleFunc("", controller.GetInvoice).Methods("GET")

	return r

}
