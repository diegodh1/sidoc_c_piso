package routes

import (
	handler "sidoc/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CreateUser func
func CreateUser(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var user handler.AppUser
		err := c.BindJSON(&user)
		switch {
		case err != nil:
			c.JSON(400, gin.H{
				"payload": nil, "message": "petición mal estructurada", "status": 400,
			})
		default:
			response := handler.CreateUser(&user, db)
			c.JSON(400, gin.H{
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
		response := handler.GetUsersERP(db)
		c.JSON(400, gin.H{
			"payload": response.Payload,
			"message": response.Message,
			"status":  response.Status,
		})
	}
	return gin.HandlerFunc(fn)
}
