package products

import "github.com/laurianderson/bootcamp_go_api_products/internal/domain"

//builder
func NewRepositoryLocal(db []*domain.Product, lastId int) Repository {
	return &repositoryLocal{
		db: db,
		lastId: lastId,
	}  
}


type repositoryLocal struct {
	db  []*domain.Product
	lastId int
}


func (rp *repositoryLocal) GetById(id int) (pr *domain.Product, err error) {
	for _, p := range rp.db {
		if p.ID == id {
                pr = p
                return
            }
		
	}
	err = ErrRepoNotFound
	return
}