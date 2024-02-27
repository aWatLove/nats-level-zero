package service

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
)

func (s *Service) PutOrderDB(order model.Order) error {
	return s.Order.Create(order)
}

func (s *Service) PutOrderCache(order model.Order) {
	s.OrderCache.PutOrder(order)
}

func (s *Service) GetFromDB(uid string) (model.Order, error) {
	return s.Order.Get(uid)
}

func (s *Service) GetFromCache(uid string) (model.Order, error) {
	return s.OrderCache.Get(uid)
}

func (s *Service) GetAllFromDB() ([]model.Order, error) {
	return s.Order.GetAll()
}

func (s *Service) GetAllFromCache() []model.Order {
	return s.OrderCache.GetAll()
}

func (s *Service) PutOrdersDBtoCache() error {
	//TODO implement me
	panic("implement me")
}
