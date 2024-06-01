package gallery_db

import (
	"database/sql"
	"errors"

	"go-mercury/pkg/constant"
)

func (d *DB) CreateProduct(product Product) (int, error) {
	db, err := d.getConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO products (title, image_url, description)
		VALUES ($1, $2, $3)
		RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, product.Title, product.ImageUrl, product.Description).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *DB) GetProductByID(id int64) (*Product, error) {
	db, err := d.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStatement := `SELECT id, title, image_url, description FROM products WHERE id = $1`
	var product Product
	err = db.QueryRow(sqlStatement, id).Scan(&product.Id, &product.Title, &product.ImageUrl, &product.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &product, constant.ProductNotFoundError
		}
		return &product, err
	}
	return &product, nil
}

func (d *DB) GetProducts() ([]Product, error) {
	var products []Product

	db, err := d.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStatement := `SELECT id, title, image_url, description FROM products`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Title, &product.ImageUrl, &product.Description)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (d *DB) UpdateProduct(product Product) (int, error) {
	db, err := d.getConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `
		UPDATE products
		SET title = $2, image_url = $3, description = $4
		WHERE id = $1`
	res, err := db.Exec(sqlStatement, product.Id, product.Title, product.ImageUrl, product.Description)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, constant.ProductNotFoundError
	}

	return int(count), nil
}

func (d *DB) DeleteProduct(id int64) (int, error) {
	db, err := d.getConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `
		DELETE FROM products
		WHERE id = $1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, constant.ProductNotFoundError
	}

	return int(count), nil
}
