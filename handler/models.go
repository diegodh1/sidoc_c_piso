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
	AppUserStatus   *bool     `gorm:"default:true;"`
	AppUserCdate    time.Time `gorm:"default:now();"`
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

//LoginUser struct login
type LoginUser struct {
	AppUserID       string `json:"AppUserID"`
	AppUserPassword string `json:"AppUserPassword"`
}

//User struct
type User struct {
	User     AppUser
	Profiles []AppUserProfile
}

//AppUserProfile struct
type AppUserProfile struct {
	AppProfileID         string
	AppUserID            string
	AppUserProfileStatus *bool     `gorm:"default:true;"`
	AppUserProfileCdate  time.Time `gorm:"default:now();"`
}

//AppProfile struct
type AppProfile struct {
	AppProfileID     string
	AppProfileStatus *bool     `gorm:"default:true;"`
	AppProfileCdate  time.Time `gorm:"default:now();"`
}

//VerificationData struct
type VerificationData struct {
	Email     string
	Code      string
	ExpiresAt time.Time
}

//UserPassReset struct
type UserPassReset struct {
	AppUserID       string `json:"AppUserID" validate:"required"`
	Code            string `json:"Code"`
	AppUserPassword string `json:"AppUserPassword"`
}

//UserPassChange struct
type UserPassChange struct {
	AppUserID          string `json:"AppUserID" validate:"required"`
	AppUserPasswordOld string `json:"AppUserPasswordOld" validate:"required"`
	AppUserPasswordNew string `json:"AppUserPasswordNew" validate:"required"`
}

//OrdenesCompraPendientes struct
type OrdenesCompraPendientes struct {
	Nit               string
	Proveedor         string
	F420Rowid         int
	F420IDTipoDocto   string
	F420Fecha         time.Time
	IDTerceroSolicita int
	UsuarioAprobador  string
	F420IndEstado     int
}

//OrdenesCompraItemsPendientes struct
type OrdenesCompraItemsPendientes struct {
	F420Rowid       int
	Codigo          int
	Descripcion     string
	Referencia      string
	Unidad          string
	Pedidas         float32
	Entradas        float32
	Pendientes      float32
	PrecioUnidad    float32
	IDEstadoItem    float32
	IdrowItemCompra int
	DetalleItem     string
	NotaItem        string
}

//EventosErp Struct
type EventosErp struct {
	EventoTipo    string
	EventoParam1  string
	EventoParam2  string
	EventoParam3  string
	EventoPruebas bool
}
