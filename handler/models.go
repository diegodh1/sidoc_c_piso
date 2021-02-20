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
	AppUserStatus   *bool `gorm:"default:true;"`
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
	AppUserID   string `json:"AppUserID"`
	AppUserPassword string `json:"AppUserPassword"`
}

//User struct
type User struct {
    User     AppUser
    Profiles []AppUserProfile
}

type AppUserProfile struct {
	AppProfileID string
	AppUserID string
	AppUserProfileStatus *bool	`gorm:"default:true;"`
	AppUserProfileCdate time.Time `gorm:"default:now();"`
}

//AppProfile struct
type AppProfile struct {
	AppProfileID string
	AppProfileStatus *bool `gorm:"default:true;"`
	AppProfileCdate time.Time `gorm:"default:now();"`
}