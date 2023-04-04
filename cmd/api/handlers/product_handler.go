package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_api_products/internal/products"
)

//builder
func NewControllerProduct(sv products.Service) *ControllerProduct {
	return &ControllerProduct{sv: sv}
}


type ControllerProduct struct {
	sv products.Service
}

func (ct *ControllerProduct) GetById() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		// process
		pr, err := ct.sv.GetById(id)
		if err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		// response
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": pr})
	}
}