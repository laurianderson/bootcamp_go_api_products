package products

import "github.com/laurianderson/bootcamp_go_api_products/internal/domain"

//Builder
func NewService(rp Repository) Service{
	return &service{rp: rp}
}

//Struct
type service struct{
	rp Repository
}

//Create new product
func (s *service) Create(pr *domain.Product) (err error){
    var lastId int
	lastId, err = s.rp.Create(pr)
	if err != nil {
		err = ErrServiceInternal
		return
	}
	pr.ID = lastId
	return
}

//Get all the products
func (s *service) GetAll() ([]*domain.Product, error){
	return s.rp.GetAll()
}

//Find product by id
func (s *service) GetById(id int) (pr *domain.Product, err error) {
	pr, err = s.rp.GetById(id)
	if err != nil {
		return
	}
	if pr == nil {
		err = ErrServiceNotFound
		return
	}
	return

}

//Update the product, select product by id
func (s *service) Update(id int, pr *domain.Product) (err error) {
	err = s.rp.Update(id, pr)
    if err!= nil {
        if err == ErrRepoNotFound {
			err = ErrServiceNotFound
			return
		}
		err = ErrServiceInternal
    }
    return
}

//Delete product by id
func (s *service) Delete(id int) (err error) {
	err = s.rp.Delete(id)
    if err!= nil {
        if err == ErrRepoNotFound {
            err = ErrServiceNotFound
            return
        }
        err = ErrServiceInternal
    }
    return
}

//Search product by price condition
func (s *service) SearchPriceGt(price float64) ([]*domain.Product, error) {
	sliceProductFound, err := s.rp.SearchPriceGt(price)
	if len(sliceProductFound) == 0 {
		return []*domain.Product{}, err
	}

	return sliceProductFound, nil
}