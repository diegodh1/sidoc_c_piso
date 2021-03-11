package routes

import (
	handler "sidoc/handler"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"encoding/json"
)

//CreateUser func
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user handler.User
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil, "message": "petición mal estructurada", "status": 400,
			})
		default:
			response := handler.CreateUser(&user.User, &user.Profiles, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

//GetUsersERP func
func GetUsersERP(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		param := c.Param("name")
		response := handler.GetUsersERP(param, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return gin.HandlerFunc(fn)
}

//UpdateProfileUser func
func UpdateProfileUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user handler.User
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil, "message": "petición mal estructurada", "status": 400,
			})
		default:
			response := handler.UpdateProfileUser(&user.User, &user.Profiles, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

//CreateProfile func
func CreateProfile(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var profile handler.AppProfile
		err := c.BindJSON(&profile)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil, "message": "petición mal estructurada", "status": 400,
			})
		default:
			response := handler.CreateProfile(&profile, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

//Login route
func Login(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var userLogin handler.LoginUser
		err := c.BindJSON(&userLogin)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"message": "Petición mal estructurada",
				"payload": nil,
				"status":  400,
			})
		default:
			response := handler.Login(userLogin.AppUserID, userLogin.AppUserPassword, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

//FindUserByID func
func FindUserByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		param := c.Param("userID")
		response := handler.FindUserByID(param, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return gin.HandlerFunc(fn)
}

//GetAllProfiles func
func GetAllProfiles(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		response := handler.GetAllProfiles(db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return gin.HandlerFunc(fn)
}

//GenerateResetPass func
func GenerateResetPass(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user handler.UserPassReset
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"message": "Petición mal estructurada",
				"payload": nil,
				"status":  400,
			})
		default:
			response := handler.GeneratePassResetCode(user.AppUserID, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}

	return gin.HandlerFunc(fn)
}

//ResetPass func
func ResetPass(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user handler.UserPassReset
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"message": "Petición mal estructurada",
				"payload": nil,
				"status":  400,
			})
		default:
			response := handler.ResetWithNewPass(&user, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}

	return gin.HandlerFunc(fn)
}

//ChangeUserPassword func
func ChangeUserPassword(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user handler.UserPassChange
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"message": "Petición mal estructurada",
				"payload": nil,
				"status":  400,
			})
		default:
			response := handler.ChangeUserPassword(&user, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}

	return gin.HandlerFunc(fn)
}

//PURSHASE ORDERS

//GetPendingOrdersByUser func
func GetPendingOrdersByUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userID := c.Param("userID")
		tipoDoc := c.Param("tipoDoc")
		nit := c.DefaultQuery("nit","")
		dateInit, _ := time.Parse("2006-01-02", c.Query("date_ini"))
		dateFinal, _ := time.Parse("2006-01-02", c.Query("date_fin"))
		ordenCompra, _ := strconv.Atoi(c.Query("ord_comp"))
		proveedor := c.Query("prov")
		response := handler.GetPendingOrdersByUser(userID, tipoDoc, nit, dateInit, dateFinal, ordenCompra, proveedor, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return gin.HandlerFunc(fn)
}

//GetPendingItemsByOrder func
func GetPendingItemsByOrder(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		orderID := c.Param("orderID")
		response := handler.GetPendingItemsByOrder(orderID, db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return gin.HandlerFunc(fn)
}

//ErpEvent func
func AddDetailsOrderCont(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var reqCont handler.ReqItemOrdPend
		bindForm := handler.FileFormReqOrders{}
		err := c.ShouldBind(&bindForm)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"message": "Petición mal estructurada",
				"payload": nil,
				"status":  400,
			})
		default:
			body := bindForm.Body
			json.Unmarshal([]byte(body), &reqCont)
			img := bindForm.Photo
			response := handler.AddDetailsOrderCont(reqCont.OrdenID, reqCont.AprobadorID, reqCont.TipoOrden, &reqCont.ListaItems, img, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}

	return gin.HandlerFunc(fn)
}

//GetAllProfiles func
func GetSpecialItems(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		response := handler.GetSpecialItems(db)
		c.JSON(response.Status, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return gin.HandlerFunc(fn)
}

func AddSale(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var sales handler.ComprasSolicitud
		err := c.BindJSON(&sales)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil, "message": "petición mal estructurada", "status": 400,
			})
		default:
			response := handler.AddSale(&sales, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

/*//ErpEvent func
func AddDetailsSales(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var []sales handler.ComprasDetalle
		err := c.BindJSON(&sales)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil, "message": "petición mal estructurada", "status": 400,
			})
		default:
			response := handler.AddSale(&sales, db)
			c.JSON(response.Status, gin.H{
				"payload": response.Payload,
				"message": response.Message,
				"status":  response.Status,
			})
		}
	}

	return gin.HandlerFunc(fn)
}*/