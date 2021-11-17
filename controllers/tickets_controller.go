package controllers

import (
	"encoding/json"
	"net/http"
    "log"
	"time"
	"github.com/llanos256/Api-Golang/commons"
	"github.com/llanos256/Api-Golang/models"
	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	tickets := []models.Tickets{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&tickets)
	json, _ := json.Marshal(tickets)
	commons.SendReponse(writer, http.StatusOK, json)
}



func Get(writer http.ResponseWriter, request *http.Request) {
	tickets := models.Tickets{}

	id := mux.Vars(request)["id"]

	db := commons.GetConnection()
	defer db.Close()
	db.Find(&tickets, id)

	if tickets.ID > 0 {
		json, _ := json.Marshal(tickets)
		commons.SendReponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func Insert (writer http.ResponseWriter, request *http.Request){
	
	tickets := models.Tickets{FEC_CREACION: time.Now(), FEC_ACTUALIZACION: time.Now()}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&tickets)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Create(&tickets).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(tickets)

	commons.SendReponse(writer, http.StatusCreated, json)
}

func Actualizar(writer http.ResponseWriter, request *http.Request){
	dato := models.Tickets{}
	updated_data := models.Tickets{FEC_ACTUALIZACION: time.Now()}
	
	id := mux.Vars(request)["id"]
	db := commons.GetConnection()
	defer db.Close()
	db.Find(&dato, id)
	if dato.ID > 0{
		err := json.NewDecoder(request.Body).Decode(&updated_data)
		if err != nil {
               commons.SendError(writer, http.StatusBadRequest)
			   return
		}
		db.Model(&dato).Updates(updated_data)
		j, _ := json.Marshal(dato)
		commons.SendReponse(writer, http.StatusOK, j)

	} else{
		commons.SendError(writer, http.StatusNotFound)
	}
}

func Eliminar (writer http.ResponseWriter, request *http.Request) {
	tickets := models.Tickets{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&tickets, id)

	if tickets.ID > 0 {
		db.Delete(tickets)
		commons.SendReponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}