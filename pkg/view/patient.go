package view

import (
	"clinic-management/pkg/handler"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePatient(ctx *gin.Context) {

	var data handler.PatientReq
	if err := ctx.ShouldBindJSON(&data); err != nil {
		slog.Error("error while binding request payload", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse JSON"})
		return
	}
	id, statusCode, err := handler.CreatePatient(data)
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

func ListAll(ctx *gin.Context) {
	search := ctx.Query("search")
	limit := ctx.DefaultQuery("limit", "10")
	offset := ctx.DefaultQuery("offset", "0")

	intlimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	intoffset, err := strconv.Atoi(offset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	data, statusCode, err := handler.ListAll(intlimit, intoffset, search)

	if err != nil {
		ctx.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(statusCode, gin.H{
		"data": data,
	})
}

func GetPatient(ctx *gin.Context) {
	idparam, isPresent := ctx.Params.Get("id")
	if !isPresent {
		slog.Error("record is not available")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no record found"})
		return
	}

	data, statusCode, err := handler.GetPatient(idparam)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"success": false, "error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{
		"sucess": true,
		"data":   data,
	})

}

func Update(ctx *gin.Context) {
	idparam, isPresent := ctx.Params.Get("id")
	if !isPresent {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "id in path is required",
		})
		return
	}

	var update handler.UpdatePatientReq
	if err := ctx.ShouldBindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "failed",
		})
		return
	}
	statusCode, err := handler.Update(idparam, update)
	if err != nil {
		ctx.JSON(statusCode, gin.H{"success": false, "error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": "Patient updated successfully",
	})
}
func Delete(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id in path is required",
		})
		return
	}

	statusCode, err := handler.Delete(id)
	if err != nil {
		ctx.JSON(statusCode, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": "Patient deleted successfully",
	})
}

