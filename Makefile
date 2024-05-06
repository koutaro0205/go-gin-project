run:
	docker-compose exec app go run main.go
start:
	docker-compose up -d
stop:
	docker-compose down
