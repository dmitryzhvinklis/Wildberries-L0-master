package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"wildberries/Nats"
	server "wildberries/Server"
	"wildberries/config"
	"wildberries/storage"
)

func main() {

	storage.CashOrders = make(map[string][]byte)
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatal("error during config downloading: ", err)
	}

	db, err := storage.ConnectToDb(conf)
	if err != nil {
		log.Fatal(err)
	}
	err = storage.CacheUP(db)
	if err != nil {
		log.Fatal(err)
	}
	sc, sub, err := Nats.GetSub(conf, db)
	if err != nil {
		log.Println(err)
	}

	go func() {
		err = server.ServerLaunch(conf)
		if err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("app Shutting down")
	if err := db.Close(); err != nil {
		log.Println(err)
	}
	if sc != nil {
		if err = sc.Close(); err != nil {
			log.Println(err)
		}
	}
	if sub != nil {
		if err = sub.Unsubscribe(); err != nil {
			log.Println(err)
		}
	}
}
