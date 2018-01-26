package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func runRestServer() {
	//run API
	log.Println("server running")
	log.Print("port: ")
	log.Println(config.RESTPort)
	router := NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":"+config.RESTPort, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
