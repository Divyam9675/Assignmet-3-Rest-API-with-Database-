package entities

import "fmt"

type Customer struct {
	Store_id   int64 `json:"Store_id"`
	Product_id       int64  `json:"Product_id"`
	Store_name string  `json:"name"`
	Phone     string   `json:"phone"`
	City      string   `json:"city"`
	Product_Price float64 `json:"Product_Price"`
	Product_name string   `json:"product_name"`

}

func (custom Customer) ToString() string {

	return fmt.Sprintf("Store_id : %d\n Product_id: %d\n Store_name: %s\n Phone: %s\n City : %s\n Product_Price : %f\n product_name : %s\n", custom.Store_id, custom.Product_id, custom.Store_name, custom.Phone, custom.City, custom.Product_Price, custom.Product_name)

}
