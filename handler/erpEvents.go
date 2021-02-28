package handler

import (
	"gorm.io/gorm"
	"strings"
)

//CreateUser func
func ErpEvent(event *EventosErp, db *gorm.DB) Response {
	switch{
		case strings.TrimSpace(event.EventoTipo) == "":
			return Response{Payload: nil, Message: "El tipo de evento es obligatorio", Status: 400}
		case strings.TrimSpace(event.EventoParam1) == "":
			return Response{Payload: nil, Message: "La orden de compra es obligatoria", Status: 400}
		case strings.TrimSpace(event.EventoParam2) == "":
			return Response{Payload: nil, Message: "El usuario ERP es obligatorio", Status: 400}		 
	}
	if err := db.Create(event).Error; err != nil {
		if strings.Contains(err.Error(), "PRIMARY KEY") {
			return Response{Payload: nil, Message: "El registro ya existe en el sistema", Status: 400}
		}
		return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
	}
	return Response{Payload: nil, Message: "Registro Realizado!", Status: 201}

}