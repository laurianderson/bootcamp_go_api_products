package handlers
/*
import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/laurianderson/bootcamp_go_api_products/cmd/api/handlers"
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
	"github.com/laurianderson/bootcamp_go_api_products/internal/products"
	"github.com/laurianderson/bootcamp_go_api_products/pkg/store"
	"github.com/stretchr/testify/assert"
)

type response struct {
	Data interface{} `json:"data"`
}

//instance server
func createServerForTestProductsHandler() *gin.Engine{
	//We do this to avoid loading unnecessary middleware
	server := gin.New()
	//to avoid innecesary logs
	gin.SetMode(gin.TestMode)

	//instances
	db, err := store.ConnectDB(os.Getenv("DB_FILE"))
		if err != nil {
			log.Fatal("Error: error connecting", err)
	}
	
	rp := products.NewRepositoryLocal(db, 0)
	s  := products.NewService(rp)
	ct := handlers.NewControllerProduct(s)

	//create routes
	server.GET("products/:product_id", ct.GetById())


	return server
}
func TestProductHandler_GetProductById(t *testing.T) {
	t.Run("should return a product", func(t *testing.T) {
		//Arrange
		var (
			expectedStatusCode = http.StatusOK
			expectedHeaders = http.Header{
				"Content-Type": []string{
					"application/json",
					"charset=utf-8",
				},
			}
			expectedResponse = response{Data: domain.Product{
				ID:          1,
				Name:        "Oil - Margarine",
				Quantity:    439,
				Code_Value:   "S82254D",
				Is_Published: true,
				Expiration:  "15/12/2021",
				Price:       71.42,
			}}
			
		)

		//Act
		server := createServerForTestProductsHandler()
        request := httptest.NewRequest(http.MethodGet, "/products/1", nil)
        request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		expectedResponse.Data = db[0]
		actual := map[string]domain.Product{}

        //Assert
        assert.Equal(t, expectedStatusCode, response.Code)
		assert.Equal(t, expectedHeaders, response.Header())
		assert.Equal(t, expectedResponse.Data, actual["data"])



	})
	t.Run("should return an error if the product is not found", func(t *testing.T) {
		//Arrange

		//Act

		//Assert
	})
	t.Run("should return an unexpected error if the service fails", func(t *testing.T) {
		//Arrange

		//Act

		//Assert
    })
}
*/