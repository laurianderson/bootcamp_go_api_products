package store

import (
	"encoding/json"
	"os"
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
)

//connect to database
func ConnectDB(filename string) ([]*domain.Product, error) {
	var products []*domain.Product
	// reader
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// decoder
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&products); err != nil {
		return nil, err
	}
	
	return products, nil
}
