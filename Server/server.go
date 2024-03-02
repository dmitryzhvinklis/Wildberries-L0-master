package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"wildberries/config"
	"wildberries/storage"
)

func handler(w http.ResponseWriter, r *http.Request) {

	html, err := ioutil.ReadFile("Templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintln(w, string(html))
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.ParseFiles("Templates/order.html"))
	var str storage.StructJsonWb

	key := r.FormValue("q")

	if key != "" {
		res, ok := storage.CashOrders[key]
		if ok {
			err = json.Unmarshal(res, &str)
			if err != nil {
				log.Println(err)
			}
			err = tmpl.Execute(w, str)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func ServerLaunch(conf *config.Config) error {
	http.HandleFunc("/", handler)

	log.Println("starting server at :8080")
	err := http.ListenAndServe(conf.Port, nil)
	if err != nil {
		return err
	}
	return nil
}
