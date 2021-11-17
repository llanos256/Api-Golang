package main

import (
	"log"
	"net/http"

	"github.com/llanos256/Api-Golang/commons"
	"github.com/llanos256/Api-Golang/routes"
	"github.com/gorilla/mux"
)

func main() {
    commons.Migrate()
	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)
	commons.EnableCORS(router)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Println("puerto de ejecución: 8000")
	log.Println(server.ListenAndServe())
}