package service

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"github.com/aWatLove/nats-lvl-zero/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

func (o OrderService) PutOrderDB(order model.Order) error {
	return o.repo.Create(order)
}

func (o OrderService) PutOrderCache(order model.Order) {
	//TODO implement me
	panic("implement me")
}

func (o OrderService) GetFromDB(uid string) (model.Order, error) {
	return o.repo.Get(uid)
}

func (o OrderService) GetFromCache(uid string) (model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderService) GetAllFromDB() ([]model.Order, error) {
	return o.repo.GetAll()
}

func (o OrderService) GetAllFromCache() ([]model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderService) PutOrdersDBtoCache() error {
	//TODO implement me
	panic("implement me")
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}
