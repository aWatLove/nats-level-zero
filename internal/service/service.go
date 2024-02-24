package service

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"github.com/aWatLove/nats-lvl-zero/internal/repository"
)

type Order interface {
	Create(order model.Order) error
	Get(uid string) (model.Order, error)
	GetAll() ([]model.Order, error)
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
