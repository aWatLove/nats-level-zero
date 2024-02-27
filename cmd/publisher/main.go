package main

import (
	"github.com/nats-io/stan.go"
	"io"
	"log"
	"os"
)

// путь до тестового JSON
var jsonPath = "static/model.json"

func main() {
	// connect to the nats streaming server
	sc, err := stan.Connect(
		"test-cluster",
		"pubID",
		stan.NatsURL("nats://0.0.0.0:4222"))
	if err != nil {
		log.Fatalf("error while connnecting to the nats streaming server: %s", err.Error())
	}
	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			log.Fatalf("error while closing publisher connection to the nats streaming server: %s", err.Error())
		}
	}(sc)
	log.Print("successfully connected to the nats streaming server")

	// parse static json file
	dataJson, err := os.Open(jsonPath)
	if err != nil {
		log.Fatalf("error while opening json file: %s", err.Error())
	}
	defer func(dataJson *os.File) {
		err = dataJson.Close()
		if err != nil {
			log.Fatalf("error while closing json fie: %s", err.Error())
		}
	}(dataJson)
	byteValue, _ := io.ReadAll(dataJson)

	// send json static repository to the nats streaming server
	err = sc.Publish("order", byteValue)
	if err != nil {
		log.Fatalf("error while publishing json static file to the nats streaming server: %s", err.Error())
	}
	log.Print("successfully published json static file to the nats streaming server")
}
