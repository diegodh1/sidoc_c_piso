package handler

import "gorm.io/gorm"

//GetSpecialItems FUNC
func GetSpecialItems(db *gorm.DB) Response {
	itemEsp := []ComprasItemEspecial{}
	db.Find(&itemEsp).Order("id_item_esp_compra")

	return Response{Payload: itemEsp, Message: "OK", Status: 200}
}


func AddSale(sale *ComprasSolicitud,db *gorm.DB) Response {
	var id uint
	db.Raw("INSERT INTO compras_solicitud(sol_compra_t_doc,sol_compra_user,sol_compra_co,sol_compra_docref,sol_compra_notas,sol_compra_estado,sol_compra_especial,sol_compra_aprobado) VALUES(?,?,?,?,?,?,?,?)",
			sale.SolCompraTDoc,
			sale.SolCompraUser,
			sale.SolCompraCo,
			sale.SolCompraDocref,
			sale.SolCompraNotas,
			sale.SolCompraEstado,
			sale.SolCompraEspecial,
			sale.SolCompraAprobado).Scan(&id)
	if id == 0 {
		return Response{Payload: nil, Message: "No se pudo crear el registro", Status: 500}
	}
		
	
	return Response{Payload: id, Message: "Registro Realizado!", Status: 201}
}

