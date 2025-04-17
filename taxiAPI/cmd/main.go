package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"awesomeProject/taxiAPI/internal/endpoints"
	"awesomeProject/taxiAPI/internal/service"
	"awesomeProject/taxiAPI/internal/storage"
)

func main() {
	rideStore := storage.NewRideMemory()
	rideService := service.NewRideService(rideStore)
	rideHandler := endpoints.NewRideHandler(rideService)
	router := mux.NewRouter()
	router.HandleFunc("/rides", rideHandler.CreateRide).Methods("POST")
	router.HandleFunc("/rides", rideHandler.GetAllRides).Methods("GET")
	router.HandleFunc("/rides/{id}", rideHandler.GetRide).Methods("GET")
	router.HandleFunc("/rides/{id}/driver", rideHandler.AssignDriverToRide).Methods("PUT")

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
