package products

import "github.com/laurianderson/bootcamp_go_api_products/internal/domain"

//builder
func NewRepositoryLocal(db []*domain.Product, lastId int) Repository {
	return &repositoryLocal{
		db: db,
		lastId: lastId,
	}  
}

//Struct
type repositoryLocal struct {
	db  []*domain.Product
	lastId int
}

//Create new product
func (rp *repositoryLocal) Create(pr *domain.Product) (lastId int, err error) {
	// set id
    rp.lastId++
	pr.ID = rp.lastId

	//append to db
	rp.db = append(rp.db, pr)

	//return lastId
	lastId = rp.lastId
    return
}

//Get all the products
func (rp *repositoryLocal) GetAll() ([]*domain.Product, error) {
	return rp.db, nil
}

//Find product by id
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

//Update the product, select product by id
func (rp *repositoryLocal) Update(id int, pr *domain.Product) (err error) {
	for i, p := range rp.db {
        if p.ID == id {
            rp.db[i] = pr
            return
        }
    }
    err = ErrRepoNotFound
    return
}

//Delete product by id
func (rp *repositoryLocal) Delete(id int) (err error) {
	for i, p := range rp.db {
        if p.ID == id {
			//consultar esto no entiendo!! Es porque se borran y se van moviendo los id (autoincrementales)
            rp.db = append(rp.db[:i], rp.db[i+1:]...)
            return
        }
    }
    err = ErrRepoNotFound
    return
}