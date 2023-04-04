package products

import( 
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
	"errors"
)

type Service interface{
	GetId(id int) (pr *domain.Product , err error)
}

var (
	ErrServiceInternal = errors.New("internal error")
	ErrServiceInvalid = errors.New("invalid product")
	ErrServiceNotUnique = errors.New("product already exists")
	ErrServiceNotFound = errors.New("product not found")
)