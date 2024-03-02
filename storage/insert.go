package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"wildberries/config"
)

func ConnectToDb(conf *config.Config) (*sql.DB, error) {

	db, err := sql.Open(conf.DriverName, conf.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Insert(uid, jsonOrder string, db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO orders_table (uid, json_order) VALUES ($1, $2 )", uid, jsonOrder)
	if err != nil {
		return err
	}
	return nil
}
