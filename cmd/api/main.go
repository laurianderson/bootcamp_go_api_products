package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_api_products/cmd/api/handlers"
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
	"github.com/laurianderson/bootcamp_go_api_products/internal/products"
)

func main() {
	//instances
	db := []*domain.Product{}
	rp := products.NewRepositoryLocal(db, 0)
	s  := products.NewService(rp)
	ct := handlers.NewControllerProduct(s)

	// server
	sv := gin.Default()
	// -> ping
	sv.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	
	prGroup := sv.Group("/products")
	{
        prGroup.GET("/:id", ct.GetById())
	}
}