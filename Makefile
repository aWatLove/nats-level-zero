build:
	docker-compose up --build

start:
	docker-compose up

test:
	go test -v ./...

#publish:
#	go run ./cmd/publisher/main.go