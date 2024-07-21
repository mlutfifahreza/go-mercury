package gallery_db

import (
	"database/sql"
	"errors"
	"fmt"

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

func (d *DB) GetProductDetail(id int64) (*ProductDetail, error) {
	db, err := d.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Query for the product
	productDetail := &ProductDetail{}
	productQuery := `SELECT id, title, image_url, description FROM products WHERE id = $1`
	err = db.QueryRow(productQuery, id).Scan(
		&productDetail.ID,
		&productDetail.Title,
		&productDetail.ImageURL,
		&productDetail.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, constant.ProductNotFoundError
		}
		return nil, err
	}

	// Query for the associated links
	linkDetailsQuery := `SELECT links.id, links.link, stores.name, stores.icon, stores.color
       FROM links 
       JOIN stores ON links.store_id = stores.id
       WHERE product_id = $1`
	rows, err := db.Query(linkDetailsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var linkDetail LinkDetail
		if err := rows.Scan(
			&linkDetail.ID,
			&linkDetail.Link,
			&linkDetail.StoreName,
			&linkDetail.StoreIcon,
			&linkDetail.StoreColor,
		); err != nil {
			return nil, err
		}
		productDetail.LinkDetails = append(productDetail.LinkDetails, linkDetail)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productDetail, nil
}

func (d *DB) GetProducts(filter ProductListFilter) ([]Product, int, error) {
	var products []Product
	total := 0

	db, err := d.getConnection()
	if err != nil {
		return nil, total, err
	}
	defer db.Close()

	sqlStatement := fmt.Sprintf(
		`SELECT id, title, image_url, description 
		FROM products 
		ORDER BY %s %s 
		OFFSET $1 
		LIMIT $2`,
		filter.OrderByField,
		filter.OrderByValue)
	rows, err := db.Query(sqlStatement, filter.Offset, filter.Limit)
	if err != nil {
		return nil, total, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Title, &product.ImageUrl, &product.Description)
		if err != nil {
			return nil, total, err
		}
		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, total, err
	}

	sqlStatement = `SELECT COUNT(*) FROM products`
	err = db.QueryRow(sqlStatement).Scan(&total)
	if err != nil {
		return nil, total, err
	}

	return products, total, nil
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
