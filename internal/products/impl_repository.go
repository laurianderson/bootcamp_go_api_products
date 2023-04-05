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

func (rp *repositoryLocal) GetAll() ([]*domain.Product, error) {
	return rp.db, nil
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