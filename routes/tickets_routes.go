package routes

import (
	"github.com/llanos256/Api-Golang/controllers"
	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/datos", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/insertar", controllers.Insert).Methods("POST")
	subRoute.HandleFunc("/eliminar/{id}", controllers.Eliminar).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
	subRoute.HandleFunc("/actualizar/{id}", controllers.Actualizar).Methods("PUT")
}