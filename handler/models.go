package handler

import (
	"time"
)

//AppUser struct
type AppUser struct {
	AppUserID       string
	AppUserName     string
	AppUserLastName string
	AppUserEmail    string
	AppUserPassword string
	AppUserErpID    int
	AppUserStatus   *bool
	AppUserCdate    time.Time
}

//UsuariosErp struct
type UsuariosErp struct {
	F552Rowid  int
	F552Nombre string
}

//Response struct
type Response struct {
	Payload interface{}
	Message string
	Status  int
}
