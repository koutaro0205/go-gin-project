run:
	docker-compose exec app go run server.go

test:
	docker-compose exec app go test ./test/... -v

test-srv:
	docker-compose exec app go test ./graph/service/... -v

test-coverage:
	docker-compose exec app go test ./test/... -v -cover -coverprofile=coverage.out -covermode=set

generate:
	docker-compose exec app go run github.com/99designs/gqlgen generate

start:
	docker-compose up -d

stop:
	docker-compose down

