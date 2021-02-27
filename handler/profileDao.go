package handler

import (
	"strings"

	"gorm.io/gorm"
)

//CreateProfile func
func CreateProfile(profile *AppProfile, db *gorm.DB) Response {
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

func assignProfile(profile *AppUserProfile, db *gorm.DB) error {
	err := db.Create(profile).Error
	return err
}

func getProfiles(userID string, db *gorm.DB) []AppUserProfile {
	profiles := []AppUserProfile{}
	db.Where("app_user_id = ?", userID).Find(&profiles)
	return profiles
}

//GetAllProfiles func
func GetAllProfiles(db *gorm.DB) Response {
	profiles := []AppUserProfile{}
	if err := db.Where("app_user_profile_status = ?", true).Find(&profiles).Error; err != nil {
		return Response{Payload: profiles, Message: "No se encontraron registros", Status: 200}
	}
	return Response{Payload: profiles, Message: "OK", Status: 200}
}

func updateAssign(profile *AppUserProfile, db *gorm.DB) error {
	profiles := []AppUserProfile{}
	if err := db.Find(&profiles, "app_profile_id = ? AND app_user_id = ?", profile.AppProfileID, profile.AppUserID).Error; err != nil || len(profiles) == 0 {
		return assignProfile(profile, db)
	}

	err := db.Where("app_user_id = ? AND app_profile_id = ?", profile.AppUserID, profile.AppProfileID).Omit("AppUserID", "AppProfileID", "AppUserProfileCdate").Updates(profile).Error

	return err
}
