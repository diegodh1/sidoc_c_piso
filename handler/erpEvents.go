package handler

import (
	"gorm.io/gorm"
)

//CreateUser func
func ErpEvent(event *EventosErp, db *gorm.DB) error {

	if err := db.Create(event).Error; err != nil {
		return err
	}

	return nil

}