package service

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"github.com/aWatLove/nats-lvl-zero/internal/repository"
)

type Order interface {
	PutOrderDB(order model.Order) error
	PutOrderCache(order model.Order)
	GetFromDB(uid string) (model.Order, error)
	GetFromCache(uid string) model.Order
	GetAllFromDB() ([]model.Order, error)
	GetAllFromCache() ([]model.Order, error)
	PutOrdersDBtoCache() error
}

type Service struct {
	repository.Order
	repository.OrderCache
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order:      repos.Order,
		OrderCache: repos.OrderCache,
	}
}
