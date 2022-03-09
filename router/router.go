package router

import (
	"GuestLedgerBookApi/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/guests", middleware.GetAllGuests).Methods("GET", "OPTIONS")
	return router
}
