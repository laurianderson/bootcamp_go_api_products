package products

import( 
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
	"errors"
)
type Repository interface {
//crear método para leer el json que usaremos como db temporalmente
//	ReadJsonDB() []domain.Product
	Create(pr *domain.Product) (lastId int, err error)
	GetById(id int) (pr *domain.Product, err error)
}

var(
	ErrRepoInternal = errors.New("internal error")
	ErrRepoNotUnique = errors.New("product already exists")
	ErrRepoNotFound = errors.New("product not found")
)
