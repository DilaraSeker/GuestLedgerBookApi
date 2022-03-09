package router

import (
	"GuestLedgerBookApi/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/guests", middleware.GetAllGuests).Methods("GET", "OPTIONS")
	router.HandleFunc("/guests", middleware.CreateGuest).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteGuest/{id}", middleware.DeleteGuest).Methods("DELETE", "OPTIONS")
	return router
}
