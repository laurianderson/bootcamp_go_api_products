package products

import( 
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
	"errors"
)

type Service interface{
	Create(pr *domain.Product) (err error)
	GetAll() ([]*domain.Product, error)
	GetById(id int) (pr *domain.Product , err error)
	Update(id int, pr *domain.Product) (err error)
	Delete(id int) (err error)
	SearchPriceGt(price float64) ([]*domain.Product, error)
}

var (
	ErrServiceInternal = errors.New("internal error")
	ErrServiceInvalid = errors.New("invalid product")
	ErrServiceNotUnique = errors.New("product already exists")
	ErrServiceNotFound = errors.New("product not found")
)