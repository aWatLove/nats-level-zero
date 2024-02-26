package model

import "time"

type Order struct {
	OrderUid          string    `json:"order_uid" gorm:"primaryKey; unique"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery" gorm:"foreignKey:OrderRefer; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payment           Payment   `json:"payment" gorm:"foreignKey:OrderRefer; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items             []Item    `json:"items" gorm:"foreignKey:OrderRefer; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created" format:"2021-11-26T06:22:19Z"`
	OofShard          string    `json:"oof_shard"`
}
