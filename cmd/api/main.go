package main

import (
	//"encoding/json"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/laurianderson/bootcamp_go_api_products/cmd/api/handlers"
	"github.com/laurianderson/bootcamp_go_api_products/cmd/api/middleware"
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
	"github.com/laurianderson/bootcamp_go_api_products/internal/products"
)

func main() {
	// .env
	if err := godotenv.Load(("/Users/landerson/Documents/meli_bootcamp/go/bootcamp_go_api_products/config.env")); err != nil {
        log.Fatal("Error: error trying to load .env file", err)
	}

/*
	//TRABAJAR ACÁ PARA RESOLVER EL USO DEL products.json
	//MODIFICAR: 1) instance db eso no va más y modificamos directamente la var rp
	//2) impl_repository.go
	//descomentar los import
	/*
	db, err := connectDB(os.Getenv("DB_FILE"))
	if err != nil {
		log.Fatal("Error: error connecting")
	}
	*/

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

		prGroup.GET("", ct.GetAll())
        prGroup.GET("/:id", ct.GetById())

		//indicate that the routes that are below the use() before continuing with the function must go through the middleware
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

//connect to database
/*
func connectDB(filename string) ([]domain.Product, error) {
	var products []domain.Product
	// reader
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// decoder
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&products); err != nil {
		return nil, err
	}
	
	return products, nil
}
*/