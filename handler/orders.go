package handler

import "gorm.io/gorm"

//GetPendingOrdersByUser func
func GetPendingOrdersByUser(userID string, tipoDoc string, db *gorm.DB) Response {
	orders := []OrdenesCompraPendientes{}
	db.Where("usuario_aprobador = ? and f420_ind_estado in (1,2) and f420_id_tipo_docto like ?", userID, tipoDoc+"%").Find(&orders)
	return Response{Payload: orders, Message: "ok", Status: 200}
}

//GetPendingItemsByOrder func
func GetPendingItemsByOrder(orderID string, db *gorm.DB) Response {
	items := []OrdenesCompraItemsPendientes{}
	db.Where("f420_rowid = ?", orderID).Find(&items)
	return Response{Payload: items, Message: "ok", Status: 200}
}
