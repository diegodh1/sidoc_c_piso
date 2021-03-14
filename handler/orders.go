package handler

import "gorm.io/gorm"
import "time"
import "mime/multipart"
import "strings"
import "strconv"
import "errors"

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
		tx = tx.Where("f420_consec_docto = ?", +ordenCompra)
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
func AddDetailsOrderCont(orderID int, aprobID string, tipo string, listaItems *[]ItemsOrdenesPendientes, db *gorm.DB) Response {
	var sucess bool
	sucess = true
	if len(*listaItems) == 0 {
		return Response{Payload: nil, Message: "No ha seleccionado ningún item para aprobar", Status: 400}		
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if tipo == "OS" {
			for _, v := range *listaItems {
				v.UsuarioAprobador = aprobID
				if err := tx.Create(&v).Error; err != nil {
					sucess = false
					tx.Rollback()
					return err
				}
			}
		}else{
			for _, v := range *listaItems {
				v.UsuarioAprobador = aprobID
				if v.Entradas > v.Pedidas {
					errors.New("Las cantidades entradas no pueden superar a las pedidas")
				}
				v.Pendientes = v.Pedidas - v.Entradas
				if err := tx.Create(&v).Error; err != nil {
					sucess = false
					tx.Rollback()
					return err
				}
			}
		}
		id := strconv.Itoa(orderID)
		event := EventosErp{EventoTipo: "EA", EventoParam1: id, EventoParam2: aprobID, EventoParam3: tipo, EventoPruebas: true}
		if err := ErpEvent(&event, db); err != nil {
				sucess = false
				tx.Rollback()
				return err
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

func AddPhotoToOrder(img *multipart.FileHeader, db *gorm.DB) Response {
	filename, err := UploadToAWS(img)
	if err != nil {
		return Response{Payload: nil, Message: "No se pudo subir la foto", Status: 500}
	}

	return Response{Payload: filename, Message: "Registro Realizado!", Status: 201}
}