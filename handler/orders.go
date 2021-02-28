package handler

import "gorm.io/gorm"
import "time"

//GetPendingOrdersByUser func
func GetPendingOrdersByUser(userID string, tipoDoc string, nit string, dateInit time.Time, dateFinal time.Time, ordenCompra int, db *gorm.DB) Response {
	orders := []OrdenesCompraPendientes{}
	tx := db.Model(&OrdenesCompraPendientes{}).Where("usuario_aprobador = ? and f420_ind_estado in (1,2) and f420_id_tipo_docto like ?", userID, tipoDoc+"%")
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
func AddDetailsOrderCont(orderID int, aprobID string, listaItems *[]ItemsCont, db *gorm.DB) Response {
	toSave := ItemsOrdenesPendientes{}
	var sucess bool
	sucess = true
	db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		for _, v := range *listaItems {
			toSave.CodCompra = orderID
			toSave.CodItem = v.ItemID
			toSave.TipoOrden = "SO"
			toSave.Entradas = v.Cantidad
			toSave.Pendiente = true
			toSave.UsuarioAprobador = aprobID
			if err := tx.Create(&toSave).Error; err != nil {
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

	return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}

}
