package models

import "time"

type OrderData struct {
	OrderUID          string    `json:"order_uid" db:"order_uid" example:"b563feb7b2b84b6test"`
	TrackNumber       string    `json:"track_number" db:"track_number" example:"WBILMTESTTRACK"`
	Entry             string    `json:"entry" db:"entry" example:"WBIL"`
	Locale            string    `json:"locale" db:"locale" example:"en"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature" example:""`
	CustomerID        string    `json:"customer_id" db:"customer_id" example:"test"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service" example:"meest"`
	ShardKey          string    `json:"shardkey" db:"shardkey" example:"9"`
	SmID              int       `json:"sm_id" db:"sm_id" example:"99"`
	DateCreated       time.Time `json:"date_created" db:"date_created" example:"2021-11-26T06:22:19Z"`
	OofShard          string    `json:"oof_shard" db:"oof_shard" example:"1"`

	Delivery DeliveryData `json:"delivery" db:"-"`
	Payment  PaymentData  `json:"payment"  db:"-"`
	Items    []ItemData   `json:"items"    db:"-"`
}

type DeliveryData struct {
	Name    string `json:"name" db:"name" example:"Test Testov"`
	Phone   string `json:"phone" db:"phone" example:"+9720000000"`
	Zip     string `json:"zip" db:"zip" example:"2639809"`
	City    string `json:"city" db:"city" example:"Kiryat Mozkin"`
	Address string `json:"address" db:"address" example:"Ploshad Mira 15"`
	Region  string `json:"region" db:"region" example:"Kraiot"`
	Email   string `json:"email" db:"email" example:"test@gmail.com"`
}

type PaymentData struct {
	Transaction  string `json:"transaction" db:"transaction" example:"b563feb7b2b84b6test"`
	RequestID    string `json:"request_id" db:"request_id" example:""`
	Currency     string `json:"currency" db:"currency" example:"USD"`
	Provider     string `json:"provider" db:"provider" example:"wbpay"`
	Amount       int    `json:"amount" db:"amount" example:"1817"`
	PaymentDT    int64  `json:"payment_dt" db:"payment_dt" example:"1637907727"`
	Bank         string `json:"bank" db:"bank" example:"alpha"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost" example:"1500"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total" example:"317"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee" example:"0"`
}

type ItemData struct {
	ChrtID      int64  `json:"chrt_id" db:"chrt_id" example:"9934930"`
	TrackNumber string `json:"track_number" db:"track_number" example:"WBILMTESTTRACK"`
	Price       int    `json:"price" db:"price" example:"453"`
	RID         string `json:"rid" db:"rid" example:"ab4219087a764ae0btest"`
	Name        string `json:"name" db:"name" example:"Mascaras"`
	Sale        int    `json:"sale" db:"sale" example:"30"`
	Size        string `json:"size" db:"size" example:"0"`
	TotalPrice  int    `json:"total_price" db:"total_price" example:"317"`
	NmID        int64  `json:"nm_id" db:"nm_id" example:"2389212"`
	Brand       string `json:"brand" db:"brand" example:"Vivienne Sabo"`
	Status      int16  `json:"status" db:"status" example:"202"`
}
