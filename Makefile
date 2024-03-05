build:
	docker-compose up --build

start:
	docker-compose up

test:
	go test -v ./...

pub:
	go run github.com/aWatLove/nats-lvl-zero/cmd/publisher