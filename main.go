package main

import (
	"log"
	handler "sidoc/handler"
	middlewares "sidoc/middlewares"
	routes "sidoc/routes"

	"github.com/gin-gonic/gin"
)

//Main func
func main() {
	r := gin.New()

	//connect to db
	var connection *handler.Config
	connection = new(handler.Config)
	connection.Initialize()
	db, err := connection.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	r.POST("/profile/create", routes.CreateProfile(db))
	r.POST("/user/login", routes.Login(db))
	r.GET("/user/erp/:name", middlewares.TokenMiddleware(2), routes.GetUsersERP(db))
	r.GET("/user/search/:userID", routes.FindUserByID(db))
	r.GET("/profile/getAll", routes.GetAllProfiles(db))
	r.POST("/user/update", middlewares.TokenMiddleware(1), routes.UpdateProfileUser(db))
	r.POST("/user/pass/code", routes.GenerateResetPass(db))
	r.POST("/user/pass/reset", routes.ResetPass(db))
	r.POST("user/changePass", routes.ChangeUserPassword(db))
	r.POST("/user/create", middlewares.TokenMiddleware(0), routes.CreateUser(db))
	//ORDERS
	r.GET("/order/user/:userID/:tipoDoc", routes.GetPendingOrdersByUser(db))
	r.GET("/order/items/:orderID", routes.GetPendingItemsByOrder(db))
	r.Run(":3000")
}
