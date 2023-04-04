package domain

type Product struct {
	ID         	  int `json:"id"`
    Name       	  string `json:"name"`
	Quantity   	  int `json:"quantity"`
	Code_Value 	  string `json:"code_value"`
	Is_Published  bool `json:"is_published"`
	Expiration    string `json:"expiration"`
	Price         float64 `json:"price"`
}