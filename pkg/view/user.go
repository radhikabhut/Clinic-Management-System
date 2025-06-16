package view

import (
	"clinic-management/pkg/handler"
	"clinic-management/pkg/model"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	var data model.UserRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		slog.Error("error while binding request payload", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse JSON"})
		return
	}
	id, statusCode, err := handler.CreateUser(data)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, gin.H{
		"data": map[string]string{
			"id": id,
		},
	})
}
func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := handler.LoginUser(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
