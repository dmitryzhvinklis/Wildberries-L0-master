package storage

import (
	"database/sql"
)

type orders struct {
	Uid  string
	Info []byte
}

func CacheUP(db *sql.DB) error {

	rows, err := db.Query("SELECT * from orders_table")
	if err != nil {
		return err
	}

	var items []orders
	for rows.Next() {
		post := orders{}
		err = rows.Scan(&post.Uid, &post.Info)
		if err != nil {
			return err
		}
		items = append(items, post)
	}
	err = rows.Close()
	if err != nil {
		return err
	}
	for _, i := range items {
		CashOrders[i.Uid] = i.Info

	}
	return nil

}
