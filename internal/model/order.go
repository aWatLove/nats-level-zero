package model

import "time"

type Order struct {
	OrderUid          string    `json:"order_uid" gorm:"primary_key; unique"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery" gorm:"foreignKey:OrderRef"`
	Payment           Payment   `json:"payment" gorm:"foreignKey:OrderRef"`
	Items             []Item    `json:"items" gorm:"foreignKey:OrderRef"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created" format:"2007-02-03T02:21:05Z"`
	OofShard          string    `json:"oof_shard"`
}
