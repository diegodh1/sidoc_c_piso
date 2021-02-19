package handler

import (
	"strings"
	"gorm.io/gorm"
)

//CreateUser func
func CreateUser(user *AppUser, profiles *[]AppUserProfile, db *gorm.DB) Response {
	switch {
	case strings.TrimSpace(user.AppUserID) == "":
		return Response{Payload: nil, Message: "El ID de usuario es obligatorio", Status: 400}
	case strings.TrimSpace(user.AppUserName) == "":
		return Response{Payload: nil, Message: "El nombre es obligatorio", Status: 400}
	case strings.TrimSpace(user.AppUserLastName) == "":
		return Response{Payload: nil, Message: "El apellido es obligatorio", Status: 400}	
	case strings.TrimSpace(user.AppUserEmail) == "" || !strings.Contains(user.AppUserEmail, "@"):
		return Response{Payload: nil, Message: "El correo es obligatorio", Status: 400}
	case user.AppUserPassword == "":
		return Response{Payload: nil, Message: "La contraseña no puede ser vacia", Status: 400}	
	case user.AppUserErpID == -1:
		return Response{Payload: nil, Message: "Debe seleccionar un usuario de ERP", Status: 400}
	default:
		if(!erpVerification(user.AppUserErpID, db)){
			return Response{Payload: nil, Message: "No es un usuario valido de ERP", Status: 400}
		}
		user.AppUserName = strings.ToUpper(user.AppUserName)
		user.AppUserLastName = strings.ToUpper(user.AppUserLastName)
		user.AppUserPassword = HashPassword(user.AppUserPassword)
		if user.AppUserPassword == "" {
			return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
		}
		if err := db.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "PRIMARY KEY") {
				return Response{Payload: nil, Message: "El registro ya existe en el sistema", Status: 400}
			}
			return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
		}

		for _, v := range *profiles {
            assignProfile(&v, db)
        }
		return Response{Payload: nil, Message: "Registro Realizado!", Status: 201}
	}
}

//GetUsersERP FUNC
func GetUsersERP(db *gorm.DB) Response {
	usuarios := []UsuariosErp{}
	db.Find(&usuarios)
	return Response{Payload: usuarios, Message: "OK", Status: 200}
}

func UpdateProfileUser(user *AppUser, profiles *[]AppUserProfile, db *gorm.DB) Response {
	switch {
	case strings.TrimSpace(user.AppUserName) == "":
		return Response{Payload: nil, Message: "El nombre es obligatorio", Status: 400}
	case strings.TrimSpace(user.AppUserLastName) == "":
		return Response{Payload: nil, Message: "El apellido es obligatorio", Status: 400}	
	case strings.TrimSpace(user.AppUserEmail) == "" || !strings.Contains(user.AppUserEmail, "@"):
		return Response{Payload: nil, Message: "El correo es obligatorio", Status: 400}
	case user.AppUserErpID == -1:
		return Response{Payload: nil, Message: "Debe seleccionar un usuario de ERP", Status: 400}
	default:
		if(!erpVerification(user.AppUserErpID, db)){
			return Response{Payload: nil, Message: "No es un usuario valido de ERP", Status: 400}
		}
		user.AppUserName = strings.ToUpper(user.AppUserName)
		user.AppUserLastName = strings.ToUpper(user.AppUserLastName)
		if queryRes := db.Where("app_user_id = ?", user.AppUserID).Omit("AppUserID", "AppUserPassword", "AppUserCdate").Updates(&user); queryRes.Error != nil || queryRes.RowsAffected == 0 {
			return Response{Payload: nil, Message: "Error al actualizar o no se encontró el usuario", Status: 404}
		}
		for _, v := range *profiles {
            updateAssign(&v, db)
        }
		
		return Response{Payload: nil, Message: "Registro actualizado!", Status: 200}
	}	
}

func erpVerification(id int, db *gorm.DB) bool{
	erpUser := []UsuariosErp{}

	if err := db.Find(&erpUser, "f552_rowid = ?", id).Error; err != nil || len(erpUser)==0{
		return false
	}
	return true
}

//Login User
func Login(userID string, password string, db *gorm.DB) Response {
	userApp := AppUser{}
	if err := scanUser(userID, &userApp, db); err != nil {
		return Response{Payload: nil, Message: "El usuario no está registrado en la base de datos", Status: 403}
	}
	switch {
	case !CheckPasswordHash(password,userApp.AppUserPassword):
		return Response{Payload: nil, Message: "Contraseña incorrecta", Status: 401}
	case *userApp.AppUserStatus == false:
		return Response{Payload: nil, Message: "El usuario no está activo en el sistema", Status: 403}
	default:
		token, err := CreateToken(userApp.AppUserID)
		if err != nil {
			return Response{Payload: nil, Message: "Error interno del servidor", Status: 500}
		}
		profiles := getProfiles(userID, db)
		var payload struct {
			User     AppUser
			Token    string
			Profiles []AppUserProfile
		}
		//profiles of the user and token
		payload.Token = token
		payload.Profiles = profiles
		payload.User = userApp
		//return
		return Response{Payload: payload, Message: "OK", Status: 200}
	}

}

func userVerification(userID string, db *gorm.DB) bool{
	userApp := []AppUser{}

	if err := db.Find(&userApp, "app_user_id = ?", userID).Error; err != nil || len(userApp)==0{
		return false
	}
	return true

}

func scanUser(userID string, userApp *AppUser, db *gorm.DB) error {

	if err := db.Raw("SELECT * FROM app_user WHERE app_user_id = ?", userID).Scan(userApp).Error; err != nil {
		return err
	}
	return nil
}

func FindUserById(userID string, db *gorm.DB) Response {
	userApp := AppUser{}
	if err := scanUser(userID, &userApp, db); err != nil {
		return Response{Payload: nil, Message: "El usuario no está registrado en la base de datos", Status: 403}
	}
	profiles := getProfiles(userID, db)
	var payload struct {
			User     AppUser
			Profiles []AppUserProfile
	}
	payload.Profiles = profiles
	payload.User = userApp

	return Response{Payload: payload, Message: "OK", Status: 200}
}