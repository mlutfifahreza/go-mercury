package gallery_db

import (
	"database/sql"
	"errors"

	"go-mercury/pkg/constant"
)

func (d *DB) CreateStore(store Store) error {
	db, err := d.getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO stores (id, name, icon, color)
		VALUES ($1, $2, $3, $4)`
	err = db.QueryRow(sqlStatement, store.Id, store.Name, store.Icon, store.Color).Err()
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetStoreByID(id string) (*Store, error) {
	db, err := d.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStatement := `SELECT id, name, icon, color FROM stores WHERE id = $1`
	var store Store
	err = db.QueryRow(sqlStatement, id).Scan(&store.Id, &store.Name, &store.Icon, &store.Color)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &store, constant.StoreNotFoundError
		}
		return &store, err
	}
	return &store, nil
}

func (d *DB) GetStores() ([]Store, error) {
	var stores []Store

	db, err := d.getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStatement := `SELECT id, name, icon, color FROM stores`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var store Store
		err := rows.Scan(&store.Id, &store.Name, &store.Icon, &store.Color)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return stores, nil
}

func (d *DB) UpdateStore(store Store) (int, error) {
	db, err := d.getConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `
		UPDATE stores
		SET name = $2, icon = $3, color = $4
		WHERE id = $1`
	res, err := db.Exec(sqlStatement, store.Id, store.Name, store.Icon)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, constant.StoreNotFoundError
	}

	return int(count), nil
}

func (d *DB) DeleteStore(id string) (int, error) {
	db, err := d.getConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	sqlStatement := `DELETE FROM stores WHERE id = $1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, constant.StoreNotFoundError
	}

	return int(count), nil
}
