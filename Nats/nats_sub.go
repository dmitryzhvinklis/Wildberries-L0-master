package Nats

import (
	"database/sql"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"wildberries/config"
	"wildberries/storage"
)

func initConn(conf *config.Config) (stan.Conn, error) {
	sc, err := stan.Connect(conf.StanClusterId, conf.ClientId, stan.NatsURL(conf.NatsUrl))

	if err != nil {
		return nil, err
	}
	return sc, nil
}

func initSub(conf *config.Config, sc stan.Conn, db *sql.DB) (stan.Subscription, error) {

	var str storage.StructJsonWb

	sub, err := sc.Subscribe(conf.Subject, func(msg *stan.Msg) {
		log.Printf("Received a message: %s\n", string(msg.Data))
		if err := json.Unmarshal(msg.Data, &str); err != nil {
			log.Println(err)
			return
		}
		storage.CashOrders[str.OrderUid] = msg.Data
		err := storage.Insert(str.OrderUid, string(msg.Data), db)
		if err != nil {
			log.Println(err)
			return
		}
	}, stan.DurableName(conf.DurableName))

	if err != nil {
		return nil, err
	}
	return sub, nil
}

func GetSub(conf *config.Config, db *sql.DB) (stan.Conn, stan.Subscription, error) {
	sc, err := initConn(conf)
	if err != nil {
		return nil, nil, err
	}
	sub, err := initSub(conf, sc, db)
	if err != nil {
		if err = sc.Close(); err != nil {
			log.Println(err)
		}
		return nil, nil, err
	}
	return sc, sub, nil
}
