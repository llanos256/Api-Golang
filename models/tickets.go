package models

import "time"

type Tickets struct {
	ID                int64  `json:"id" gorm:"primary_key;auto_increment"`
	USUARIO           string `json:"usuario"`
	FEC_CREACION      time.Time `json:"fec_creacion"`
	FEC_ACTUALIZACION time.Time    `json:"fec_actualizacion"`
	ESTADO            string  `json:"estado"`
}