

all :
	docker-compose up

build:
	docker-compose builb

serv:
	go run cmd/main.go

example:
	go run pub_srcipt/publish.go  one.json