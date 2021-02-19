package main

import (
	"log"
	handler "sidoc/handler"
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
	r.POST("/user/create", routes.CreateUser(db))
	r.POST("/profile/create", routes.CreateProfile(db))
	r.POST("/user/login", routes.Login(db))
	r.GET("/user/erp", routes.GetUsersERP(db))
	r.GET("/user/search/:userID", routes.FindUserById(db))
	r.GET("/profile/getAll", routes.GetAllProfiles(db))
	r.POST("/user/update", routes.UpdateProfileUser(db))
	r.Run(":3000")
}
