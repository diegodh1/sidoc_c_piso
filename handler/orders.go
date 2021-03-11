package handler

import "gorm.io/gorm"
import "time"
import "mime/multipart"
import "strings"
import "strconv"

//GetPendingOrdersByUser func
func GetPendingOrdersByUser(userID string, tipoDoc string, nit string, dateInit time.Time, dateFinal time.Time, ordenCompra int, proveedor string, db *gorm.DB) Response {
	orders := []OrdenesCompraPendientes{}
	tx := db.Model(&OrdenesCompraPendientes{}).Where("usuario_aprobador = ? and f420_ind_estado in (1,2) and f420_id_tipo_docto like ?", userID, tipoDoc+"%").Order("f420_fecha desc")
	//Select("TOP(?) *", 20)
	//db.Where("usuario_aprobador = ? and f420_ind_estado in (1,2) and f420_id_tipo_docto like ?", userID, tipoDoc+"%").Find(&orders)
	if nit != "" {
		tx = tx.Where("nit like ?", nit+"%")
	}
	if ordenCompra > 0 {
		tx = tx.Where("f420_rowid = ?", +ordenCompra)
	}
	if !dateInit.IsZero() || !dateFinal.IsZero() {
		tx = tx.Where("f420_fecha BETWEEN (?) AND (?)", dateInit, dateFinal)
	}
	if proveedor != "" {
		tx = tx.Where("proveedor like ?", proveedor+"%")
	}
	tx.Find(&orders)
	return Response{Payload: orders, Message: "ok", Status: 200}
}

//GetPendingItemsByOrder func
func GetPendingItemsByOrder(orderID string, db *gorm.DB) Response {
	items := []OrdenesCompraItemsPendientes{}
	db.Where("f420_rowid = ?", orderID).Find(&items)
	return Response{Payload: items, Message: "ok", Status: 200}
}

//AddDetailsOrderCont func
func AddDetailsOrderCont(orderID int, aprobID string, tipo string, listaItems *[]ItemsOrdenesPendientes, img *multipart.FileHeader, db *gorm.DB) Response {
	var sucess bool
	sucess = true
	err := db.Transaction(func(tx *gorm.DB) error {
		id := strconv.Itoa(orderID)
		event := EventosErp{EventoTipo: tipo, EventoParam1: id, EventoParam2: aprobID, EventoPruebas: true}
		if err := ErpEvent(&event, db); err != nil {
				sucess = false
				tx.Rollback()
				return err
		}
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		for _, v := range *listaItems {
			v.CodCompra = orderID
			v.UsuarioAprobador = aprobID
			if err := tx.Create(&v).Error; err != nil {
				sucess = false
				tx.Rollback()
				return err
			}
		}
		return tx.Commit().Error
	})
	if sucess {
		return Response{Payload: nil, Message: "Registro Realizado!", Status: 201}
	}
	
	if strings.Contains(err.Error(), "PRIMARY KEY") {
		return Response{Payload: nil, Message: "Esta Orden ya fue revisada", Status: 400}
	}

	return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}

}
