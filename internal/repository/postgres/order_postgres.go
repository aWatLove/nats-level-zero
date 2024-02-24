package postgres

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"gorm.io/gorm"
)

type OrderPostgres struct {
	db *gorm.DB
}

func (o OrderPostgres) Create(order model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderPostgres) Get(uid string) (model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderPostgres) GetAll() ([]model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderPostgres(db *gorm.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}
