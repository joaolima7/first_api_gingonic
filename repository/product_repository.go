package repository

import (
	"database/sql"
	"first-api-gin/model"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Erro no banco: " + err.Error())
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println("Erro no banco: " + err.Error())
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil

}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) FindProductByID(id_product int) (*model.Product, error) {
	var product model.Product
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = query.QueryRow(id_product).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &product, nil

}

func (pr *ProductRepository) UpdateProduct(updates map[string]interface{}, id int) (bool, error) {
	if len(updates) == 0 {
		return false, fmt.Errorf("no fields to update")
	}

	query := "UPDATE product SET "
	params := []interface{}{}
	counter := 1

	for key, value := range updates {
		if counter > 1 {
			query += ", "
		}
		query += fmt.Sprintf("%s = $%d", key, counter)
		params = append(params, value)
		counter++
	}

	query += fmt.Sprintf(" WHERE id = $%d", counter)
	params = append(params, id)

	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(params...)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return rowsAffected > 0, nil
}

func (pr *ProductRepository) DeleteProduct(idProduct int) (bool, error) {
	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		return false, err
	}

	defer query.Close()

	result, err := query.Exec(idProduct)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
