package main

import (
	"GuestLedgerBookApi/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	r := router.Router()
	fmt.Println("Server started on the port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
