package cache

import (
	"errors"
	"fmt"
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"net/http"
)

type OrderCache struct {
	cch *Cache
}

func NewOrderCache(cch *Cache) *OrderCache {
	return &OrderCache{cch: cch}
}

func (o *OrderCache) PutOrder(order model.Order) {
	uid := order.OrderUid
	o.cch.Mutex.Lock()
	o.cch.Data[uid] = order
	o.cch.Mutex.Unlock()
}

func (o *OrderCache) Get(uid string) (model.Order, error) {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()
	if value, ok := o.cch.Data[uid]; ok {
		return value, nil
	}
	return model.Order{}, NewErrorHandler(errors.New(fmt.Sprintf("can't find order with uid: %s", uid)), http.StatusBadRequest)
}

func (o *OrderCache) GetAll() []model.Order {
	o.cch.Mutex.Lock()
	defer o.cch.Mutex.Unlock()
	if len(o.cch.Data) == 0 {
		return []model.Order{}
	}

	orders := make([]model.Order, len(o.cch.Data))

	i := 0
	for _, val := range o.cch.Data {
		orders[i] = val
		i++
	}

	return orders
}
