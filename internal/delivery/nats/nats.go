package nats

import (
	"encoding/json"
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"github.com/aWatLove/nats-lvl-zero/internal/service"
	"github.com/nats-io/stan.go"
	"log"
	"sync"
)

type Nats struct {
	service *service.Service
}

func NewNats(service *service.Service) *Nats {
	return &Nats{service: service}
}

func (n Nats) Connect(clusterId, clientId, natsUrl string) (stan.Conn, error) {
	conn, err := stan.Connect(clusterId, clientId, stan.NatsURL(natsUrl))
	if err != nil {
		log.Printf("error while connecting to nats streaming subject: %s", err.Error())
		return nil, err
	}
	log.Println("successfully connecting to nats streaming subject")
	return conn, nil
}

func (n Nats) Subscribe(wg *sync.WaitGroup, sc stan.Conn, natsSubject string) error {
	defer wg.Done()

	sub, err := sc.Subscribe(natsSubject, func(msg *stan.Msg) {
		order, err := n.unmarshall(msg)
		if err != nil {
			return
		}
		err = n.service.PutOrderDB(order)
		if err != nil {
			return
		}
		n.service.PutOrderCache(order)

		log.Printf("success adding order with uid: %s", order.OrderUid)
	})
	if err != nil {
		log.Printf("error while subscribing to nats-streaming subject: %s", err.Error())
		return err
	}

	for {
		if !sub.IsValid() {
			wg.Done()
			break
		}
	}

	err = sub.Unsubscribe()
	if err != nil {
		log.Printf("error while unsubscribing from nats-streaming subject: %s\n", err.Error())
		return err
	}
	log.Println("successfully unsubscribing from nats-streaming subject")
	return nil
}

func (n Nats) unmarshall(msg *stan.Msg) (model.Order, error) {
	order := model.Order{}
	err := json.Unmarshal(msg.Data, &order)

	// todo validating data

	if err != nil {
		log.Printf("error while unmarshalling msg to model: %s", err.Error())
		return order, err
	}
	return order, nil
}
