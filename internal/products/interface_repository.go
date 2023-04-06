package products

import( 
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
	"errors"
)
type Repository interface {
	//crud
	Create(pr *domain.Product) (lastId int, err error)
	GetAll() ([]*domain.Product, error)
	GetById(id int) (pr *domain.Product, err error)
	Update(id int, pr *domain.Product) (err error)
	Delete(id int) (err error)

	//add method search for price 
	SearchPriceGt(price float64) ([]*domain.Product, error)
	
}

var(
	ErrRepoInternal = errors.New("internal error")
	ErrRepoNotUnique = errors.New("product already exists")
	ErrRepoNotFound = errors.New("product not found")
)
