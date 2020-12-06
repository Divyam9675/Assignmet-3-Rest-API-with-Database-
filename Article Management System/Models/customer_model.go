package Models

import (
	"database/sql"
	"entities"
)

type CustomerModel struct {
	Db *sql.DB
}

func (custmModi CustomerModel) FindAll() (custom []entities.Customer, err error) {

	row, err := custmModi.Db.Query("Select *from articles")
	if err != nil {
		return nil, err
	} else {

		custmMod := []entities.Customer{}

		for row.Next() {

			var id int64
			var title string
			var description string
			var content string
			err2 := row.Scan(&id, &title, &description, &content)

			if err2 != nil {
				return nil, err2
			} else {

				custm := entities.Customer{

					ID:          id,
					TITLE:       title,
					DESCRIPTION: description,
					CONTENT:     content,
				}

				//var custmMod []entities.Customer

				custmMod = append(custmMod, custm)

			}

		}

		return custmMod, nil

	}
}

func (custmModi CustomerModel) Search(keyword string) (custom []entities.Customer, err error) {

	row, err := custmModi.Db.Query("Select *from articles where id like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {

		custmMod := []entities.Customer{}

		for row.Next() {

			var id int64
			var title string
			var description string
			var content string
			err2 := row.Scan(&id, &title, &description, &content)

			if err2 != nil {
				return nil, err2
			} else {

				custm := entities.Customer{

					ID:          id,
					TITLE:       title,
					DESCRIPTION: description,
					CONTENT:     content,
				}

				//var custmMod []entities.Customer

				custmMod = append(custmMod, custm)

			}

		}

		return custmMod, nil

	}
}

func (cust CustomerModel) Create(custum *entities.Customer) (err error) {

	result, err := cust.Db.Exec("insert into articles(title, description, content) values(?,?,?)", custum.TITLE, custum.DESCRIPTION, custum.CONTENT)
	if err != nil {
		return err
	} else {
		custum.ID, _ = result.LastInsertId()
		return nil

	}
}

func (cust CustomerModel) Update(custum *entities.Customer) (int64, error) {

	result, err := cust.Db.Exec("update articles set title = ?, description = ?, content = ? where id  = ? ", custum.TITLE, custum.DESCRIPTION, custum.CONTENT, custum.ID)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()

	}

}

func (cust CustomerModel) Delete(id int64) (int64, error) {

	result, err := cust.Db.Exec("delete from articles where id = ? ", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()

	}

}
