package store

import (
	"encoding/json"
	"os"
	"github.com/laurianderson/bootcamp_go_api_products/internal/domain"
)

//connect to database
func ConnectDB(filename string) ([]*domain.Product, error) {
	var products []*domain.Product
	// open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	
	// reader file
	reader := json.NewDecoder(file)
	if err := reader.Decode(&products); err != nil {
		return nil, err
	}

	//VER ESTO MODIFICAR!!!!!!!!!!!
	/*
	//writter file
	writter := json.NewEncoder(file)
	if err := writter.Encode(&products); err != nil {
		return nil, err
    }
	*/

	return products, nil
}
