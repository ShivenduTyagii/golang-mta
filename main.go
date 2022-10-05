package main

import (
	"log"
	"net/http"

	"github.com/ShivenduTyagii/GO-MTA/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/mta", controllers.GetMtaData).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", r))

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}
