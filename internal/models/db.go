package models

import "time"

type OrderDataDb struct {
	OrderUID          string    `json:"order_uid" db:"order_uid"`
	TrackNumber       string    `json:"track_number" db:"track_number"`
	Entry             string    `json:"entry" db:"entry"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerID        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	ShardKey          string    `json:"shardkey" db:"shardkey"`
	SmID              int       `json:"sm_id" db:"sm_id"`
	DateCreated       time.Time `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`

	DeliveryData // we need nesting because pgxscan.Select won’t see the fields.
	PaymentData  // and again

	Items []ItemData `json:"items"    db:"-"`
}
