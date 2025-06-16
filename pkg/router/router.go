package router

import (
	"clinic-management/pkg/middleware"
	"clinic-management/pkg/view"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.RouterGroup) {

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"status": "ok"})
	})

	patientRoute := app.Group("/patients")
	patientRoute.Use(middleware.AuthMiddleware())
	patientRoute.POST("/", middleware.RoleMiddleware("receptionist"), view.CreatePatient)
	patientRoute.GET("/", middleware.RoleMiddleware("receptionist", "doctor"), view.ListAll)
	patientRoute.GET("/:id", middleware.RoleMiddleware("receptionist", "doctor"), view.GetPatient)
	patientRoute.PUT("/:id", middleware.RoleMiddleware("receptionist", "doctor"), view.Update)
	patientRoute.DELETE("/:id", middleware.RoleMiddleware("receptionist"), view.Delete)

	userRoutes := app.Group("/users")
	userRoutes.POST("/", view.CreateUser)

	authRoutes := app.Group("/auth")
	authRoutes.POST("/login", view.Login)
}
