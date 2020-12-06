package Models

import (
	"database/sql"
	"entities"
)

type CustomerModel struct {
	Db *sql.DB
}

func (custmModi CustomerModel) FindAll() (custom []entities.Customer, err error) {

	row, err := custmModi.Db.Query("Select *from product")
	if err != nil {
		return nil, err
	} else {

		custmMod := []entities.Customer{}

		for row.Next() {

			var S_id    int64
			var P_id  int64
			var S_name string 
			var num     string 
			var cty string
			var P_Price float64
			var P_name string

			err2 := row.Scan(&S_id, &P_id, &S_name, &num, &cty, &P_Price, &P_name)

			if err2 != nil {
				return nil, err2
			} else {

				custm := entities.Customer{

					Store_id:          S_id,
					Product_id:       P_id,
					Store_name:       S_name,
					Phone:            num,
					City:             cty,
					Product_Price:    P_Price,
					Product_name:     P_name,
				}

				//var custmMod []entities.Customer

				custmMod = append(custmMod, custm)

			}

		}

		return custmMod, nil

	}
}

func (custmModi CustomerModel) Search(keyword string) (custom []entities.Customer, err error) {

	row, err := custmModi.Db.Query("Select *from product where Store_id like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {

		custmMod := []entities.Customer{}

		for row.Next() {

			var S_id    int64
			var P_id  int64
			var S_name string 
			var num     string 
			var cty string
			var P_Price float64
			var P_name string

			err2 := row.Scan(&S_id, &P_id, &S_name, &num, &cty, &P_Price, &P_name)

			if err2 != nil {
				return nil, err2
			} else {

				custm := entities.Customer{

					Store_id:          S_id,
					Product_id:       P_id,
					Store_name:       S_name,
					Phone:            num,
					City:             cty,
					Product_Price:    P_Price,
					Product_name:     P_name,
				}

				//var custmMod []entities.Customer

				custmMod = append(custmMod, custm)

			}

		}

		return custmMod, nil

	}
}

func (cust CustomerModel) Create(custum *entities.Customer, keyword int64) (err error) {

	result, err := cust.Db.Exec("insert into product(Product_id, Store_name, Phone, City, Product_Price, product_name) values(?,?,?,?,?,?)", custum.Product_id, custum.Store_name, custum.Phone, custum.City, custum.Product_Price, custum.Product_name)
	
	if err != nil {
		return err
	} else {
		custum.Store_id, _ = result.LastInsertId()
		return nil

	}
}

func (cust CustomerModel) Update(custum *entities.Customer) (int64, error) {

	result, err := cust.Db.Exec("update product set Product_id = ?, Store_name =? , Phone = ?, City = ?, Product_Price = ?, Product_name = ? where Store_id  = ? ", custum.Product_id, custum.Store_name, custum.Phone, custum.City, custum.Product_Price, custum.Product_name, custum.Store_id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()

	}

}

func (cust CustomerModel) Delete(id int64) (int64, error) {

	result, err := cust.Db.Exec("delete from product where Store_id = ? ", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()

	}

}
