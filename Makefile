run:
	docker-compose exec app go run main.go

server:
	docker-compose exec app go run server.go

generate:
	docker-compose exec app go run github.com/99designs/gqlgen generate

start:
	docker-compose up -d

stop:
	docker-compose down

