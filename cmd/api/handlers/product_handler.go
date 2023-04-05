package handlers

import (
	"errors"
	"os"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
	"github.com/laurianderson/bootcamp_go_api_products/internal/products"
)

//builder
func NewControllerProduct(sv products.Service) *ControllerProduct {
	return &ControllerProduct{sv: sv}
}


type ControllerProduct struct {
	sv products.Service
}

func(ct *ControllerProduct) Create() gin.HandlerFunc{
	type request struct {
		Name         string `json:"name" biding:"required"`
		Quantity     int `json:"quantity" biding:"required"`
		Code_Value 	 string `json:"code_value" biding:"required"`
		Is_Published bool `json:"is_published" biding:"required"`
		Expiration   string `json:"expiration" biding:"required"`
		Price        float64 `json:"price" biding:"required"`
	}
	return func(ctx *gin.Context){
		//add token in the header
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"mesagge": "invalid token"})
			return
		}
		//request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		//process
		pr := &domain.Product{
			Name: req.Name,
            Quantity: req.Quantity,
            Code_Value: req.Code_Value,
            Is_Published: req.Is_Published,
            Expiration: req.Expiration,
            Price: req.Price,
		}
		err := ct.sv.Create(pr)
		if err != nil {
			if errors.Is(err, products.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid product"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
        }
		//response
		ctx.JSON(http.StatusCreated, gin.H{"message": "success", "data": pr})
	}
}

func(ct *ControllerProduct) GetAll() gin.HandlerFunc{
	return func(ctx *gin.Context){
		//request

		//process
		prs, err := ct.sv.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get movies"})
			return
		}
        //response
        ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": prs})
    }
}

func (ct *ControllerProduct) GetById() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		//process
		pr, err := ct.sv.GetById(id)
		if err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		//response
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": pr})
	}
}

func (ct *ControllerProduct) Update() gin.HandlerFunc{
	type request struct {
		Name         string `json:"name" biding:"required"`
		Quantity     int `json:"quantity" biding:"required"`
		Code_Value 	 string `json:"code_value" biding:"required"`
		Is_Published bool `json:"is_published" biding:"required"`
		Expiration   string `json:"expiration" biding:"required"`
		Price        float64 `json:"price" biding:"required"`
	}
	return func(ctx *gin.Context) {
		//add token in the header
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"mesagge": "invalid token"})
			return
		}

		//request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		var req request 
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		//process
		pr := &domain.Product{
			ID: id,
			Name: req.Name,
            Quantity: req.Quantity,
            Code_Value: req.Code_Value,
            Is_Published: req.Is_Published,
            Expiration: req.Expiration,
            Price: req.Price,
		}
		if err := ct.sv.Update(id, pr); err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			if errors.Is(err, products.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid product"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		//response
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": pr})
	}
}

func (ct *ControllerProduct) UpdatePartial() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//add token in the header
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"mesagge": "invalid token"})
			return
		}

		//request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		// -> get product by id
		pr, err := ct.sv.GetById(id)
		if err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}
		if err := ctx.ShouldBindJSON(&pr); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		pr.ID = id

		//process
		if err := ct.sv.Update(id, pr); err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			if errors.Is(err, products.ErrServiceInvalid) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid product"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		//response
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": pr})
	}
      
}

func (ct *ControllerProduct) Delete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//add token in the header
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"mesagge": "invalid token"})
			return
		}
		
		// request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		// process
		if err := ct.sv.Delete(id); err != nil {
			if errors.Is(err, products.ErrServiceNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}

		// response
		//ctx.Header("Location", fmt.Sprintf("/movies/%d", id))
		ctx.JSON(http.StatusNoContent, nil)
	}
}