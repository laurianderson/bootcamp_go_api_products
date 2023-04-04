package products

import "github.com/laurianderson/bootcamp_go_api_products/internal/domain"

//builder
func NewService(rp Repository) Service{
	return &service{rp: rp}
}

type service struct{
	rp Repository
}

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

func (s *service) GetAll() ([]*domain.Product, error){
	return s.rp.GetAll()
}

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