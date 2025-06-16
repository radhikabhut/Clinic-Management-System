package main

import (
	"clinic-management/config"
	"clinic-management/pkg/db"
	"clinic-management/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig("F:\\workspace\\Clinic-Management\\.env")
	db.ConnectDB()

	app := gin.Default()

	v1Route := app.Group("v1")
	router.InitRoute(v1Route)

	app.Run(":8080")

}
