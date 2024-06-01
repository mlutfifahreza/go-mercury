package gallery_db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"go-mercury/pkg/constant"
)

type DB struct {
	host     string
	port     int
	username string
	password string
	dbname   string
}

func NewDB(
	host string,
	port int,
	username string,
	password string,
	dbname string) *DB {
	return &DB{
		host:     host,
		port:     port,
		username: username,
		password: password,
		dbname:   dbname,
	}
}

func (d *DB) getConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.username, d.password, d.dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v\n", err)
	}
	return db, err
}

func (d *DB) CreateProduct(product Product) (int, error) {
	db, err := d.getConnection()
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

func (d *DB) GetProductByID(id int64) (Product, error) {
	db, err := d.getConnection()
	defer db.Close()

	sqlStatement := `SELECT id, title, image_url, description FROM products WHERE id = $1`
	var product Product
	err = db.QueryRow(sqlStatement, id).Scan(&product.Id, &product.Title, &product.ImageUrl, &product.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return product, constant.ProductNotFoundError
		}
		return product, err
	}
	return product, nil
}

func (d *DB) GetProducts() ([]Product, error) {
	var products []Product

	db, err := d.getConnection()
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

	fmt.Printf("%d rows affected\n", count)
	return int(count), nil
}
