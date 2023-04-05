package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
	"github.com/laurianderson/bootcamp_go_api_products/internal/products"
)

//Builder
func NewControllerProduct(sv products.Service) *ControllerProduct {
	return &ControllerProduct{sv: sv}
}

//Struct
type ControllerProduct struct {
	sv products.Service
}

//Create new product
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
		//request
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		//validate expiration
		valid, _ := validateExpiration(req.Expiration)
		if !valid {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid expiration"})
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

//Get all the products
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

//Find product by id
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

//Update the product, select product by id
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

//Patch the product, select product by id
func (ct *ControllerProduct) UpdatePartial() gin.HandlerFunc{
	return func(ctx *gin.Context) {
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

//Delete product by id
func (ct *ControllerProduct) Delete() gin.HandlerFunc{
	return func(ctx *gin.Context) {
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

// validateExpiration confirm that the expiration date is valid
func validateExpiration(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) || (list[1] < 1 || list[1] > 12) || (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}