package handler

import (
	"time"
	"mime/multipart" 
)

//AppUser struct
type AppUser struct {
	AppUserID       string
	AppUserName     string
	AppUserLastName string
	AppUserEmail    string
	AppUserPassword string
	AppUserErpID    int
	AppUserErpName 	string
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
	EventoPruebas bool
}

//ItemsOrdenPendientes Struct
type ItemsOrdenesPendientes struct {
	CodCompra	int
	CodItem 	int
	TipoOrden string
	Referencia string
	UnidadMedida string
	Entradas float32
	Pendientes float32
	Pendiente bool
	UsuarioAprobador string
}

type ItemsCont struct {
	ItemID int
	Cantidad float32
}

//ReqItemOrdPend Struct
type ReqItemOrdPendCont struct {
	OrdenID int
	AprobadorID string
	ListaItems []ItemsCont
}

type FileFormReqOrders struct {
	Body string `form:"body" binding:"required"`
	Photo *multipart.FileHeader `form:"photo" binding:"required"`	
}

type ComprasItemEspecial struct {
	IdItemEspCompra int
	ItemEspCodigo int
	ItemEspReferencia string
	ItemEspDescripcion string
	ItemEspActivo bool
	ItemEspUsucrea string
}

type ComprasSolicitud struct {
	IdSolCompra uint `gorm:"primaryKey;autoIncrement:true"`
	SolCompraTDoc string
	SolCompraUser uint
	SolCompraCo string
	SolCompraDocref string
	SolCompraNotas string
	SolCompraEstado uint
	SolCompraEspecial bool
	SolCompraAprobado time.Time
}

type ComprasDetalle struct {
	IdSolCompra uint 
	CompraDetCodigo uint
	CompraDetDescripcion string
	CompraDetUnd string
	CompraDetCantidad float32
	CompraDetNota string
	CompraDetFechae time.Time
	CompraDetUser uint
}
