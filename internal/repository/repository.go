package repository

import "github.com/aWatLove/nats-lvl-zero/internal/model"

type Order interface {
	Create(order model.Order)
}
