package service

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"github.com/aWatLove/nats-lvl-zero/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

func (o OrderService) Create(order model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderService) Get(uid string) (model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderService) GetAll() ([]model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}
