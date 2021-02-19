package handler

import (
	"strings"
	"gorm.io/gorm"

)

func CreateProfile(profile *AppProfile, db *gorm.DB) Response{
	if strings.TrimSpace(profile.AppProfileID) == "" {
		return Response{Payload: nil, Message: "El ID de usuario es obligatorio", Status: 400}
	}
	profile.AppProfileID = strings.ToUpper(profile.AppProfileID)
	if err := db.Create(&profile).Error; err != nil {
		if strings.Contains(err.Error(), "PRIMARY KEY") {
			return Response{Payload: nil, Message: "El registro ya existe en el sistema", Status: 400}
		}
		return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
	}
	return Response{Payload: nil, Message: "Registro Realizado!", Status: 201}
}

func assignProfile(profile *AppUserProfile, db *gorm.DB) error{
	err := db.Create(profile).Error
	return err
}

func getProfiles(userID string, db *gorm.DB) []AppUserProfile {
	profiles := []AppUserProfile{}
	if err := db.Where("app_user_id = ? and app_user_profile_status = ?", userID, true).Find(&profiles).Error; err != nil {
		return profiles
	}
	return profiles
}