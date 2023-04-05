package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/laurianderson/bootcamp_go_api_products/cmd/api/handlers"
	"github.com/laurianderson/bootcamp_go_api_products/cmd/api/middleware"
	"github.com/laurianderson/bootcamp_go_api_products/internal/products"
	"github.com/laurianderson/bootcamp_go_api_products/pkg/store"
)

func main() {
	// .env
	if err := godotenv.Load(("/Users/landerson/Documents/meli_bootcamp/go/bootcamp_go_api_products/config.env")); err != nil {
        log.Fatal("Error: error trying to load .env file", err)
	}

	//instances
	db, err := store.ConnectDB(os.Getenv("DB_FILE"))
	if err != nil {
		log.Fatal("Error: error connecting", err)
	}

	rp := products.NewRepositoryLocal(db, 0)
	s  := products.NewService(rp)
	ct := handlers.NewControllerProduct(s)

	// server
	sv := gin.Default()

	// -> ping
	sv.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	//group the routes
	prGroup := sv.Group("/products")
	{
		prGroup.GET("", ct.GetAll())
        prGroup.GET("/:id", ct.GetById())

		// middleware in this routes
		prGroup.Use(middleware.MiddlewareVerificationToken())
		prGroup.POST("", ct.Create())
		prGroup.PUT("/:id", ct.Update())
		prGroup.PATCH("/:id", ct.UpdatePartial())
		prGroup.DELETE("/:id", ct.Delete())
	}

	// run
	if err := sv.Run(); err != nil {
		panic(err)
	}
}
