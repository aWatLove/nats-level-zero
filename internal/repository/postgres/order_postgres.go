package postgres

import (
	"fmt"
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"gorm.io/gorm"
	"log"
)

type OrderPostgres struct {
	db *gorm.DB
}

func (o OrderPostgres) Create(order model.Order) error {
	err := o.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			fmt.Println("order in tx", order)
			return err
		}

		return nil
	})
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
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
