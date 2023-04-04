package products

import "github.com/laurianderson/bootcamp_go_api_products/internal/domain"

//construct
func NewService(rp Repository) Service{
	return &service{rp: rp}
}

type service struct{
	rp Repository
}

func (s *service) GetId(id int) (pr *domain.Product, err error) {
	pr, err = s.rp.GetId(id)
	if err != nil {
		return
	}
	if pr == nil {
		err = ErrServiceNotFound
		return
	}
	return

}