migrate:
	migrate -database "postgres://root:root@127.0.0.1:5432/simpleshop?sslmode=disable" -path migrations up
down:
	migrate -database "postgres://root:root@127.0.0.1:5432/simpleshop?sslmode=disable" -path migrations down

build:
	mkdir -p dist
	go build -o dist cmd/main.go