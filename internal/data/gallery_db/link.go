package gallery_db

import (
	"database/sql"
	"errors"

	"go-mercury/pkg/constant"
)

func (d *DB) CreateLink(link Link) (int, error) {
	db, err := d.getConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO links (product_id, store_id, link)
		VALUES ($1, $2, $3)
		RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, link.ProductId, link.StoreId, link.Link).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *DB) GetLinkByID(id int64) (*Link, error) {
	db, err := d.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStatement := `SELECT id, product_id, store_id, link FROM links WHERE id = $1`
	var link Link
	err = db.QueryRow(sqlStatement, id).Scan(&link.Id, &link.ProductId, &link.StoreId, &link.Link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &link, constant.LinkNotFoundError
		}
		return &link, err
	}
	return &link, nil
}

func (d *DB) GetLinks() ([]Link, error) {
	var links []Link

	db, err := d.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStatement := `SELECT id, product_id, store_id, link FROM links`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var link Link
		err := rows.Scan(&link.Id, &link.ProductId, &link.StoreId, &link.Link)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return links, nil
}

func (d *DB) UpdateLink(link Link) (int, error) {
	db, err := d.getConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `
		UPDATE links
		SET product_id = $2, store_id = $3, link = $4
		WHERE id = $1`
	res, err := db.Exec(sqlStatement, link.Id, link.ProductId, link.StoreId, link.Link)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, constant.LinkNotFoundError
	}

	return int(count), nil
}

func (d *DB) DeleteLink(id int64) (int, error) {
	db, err := d.getConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `
		DELETE FROM links
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
		return 0, constant.LinkNotFoundError
	}

	return int(count), nil
}
