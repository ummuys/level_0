package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/ummuys/level_0/internal/models"
)

const (
	qInsertInfo = `
INSERT INTO orders.info (
  order_uid, track_number, entry, locale, internal_signature,
  customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
`

	qInsertPayment = `
INSERT INTO orders.payments (
  order_uid, transaction_id, request_id, currency, provider,
  amount, payment_dt, bank, delivery_cost, goods_total, custom_fee
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
`

	qInsertDelivery = `
INSERT INTO orders.deliveries (
  order_uid, name, phone, zip, city, address, region, email
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
`

	qInsertItem = `
INSERT INTO orders.items (
  order_uid, chrt_id, track_number, price, rid, name, sale, size,
  total_price, nm_id, brand, status
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
`
)

func createBatch(orderData *models.OrderData) (*pgx.Batch, int) {
	b := &pgx.Batch{}
	n := 0

	b.Queue(qInsertInfo,
		orderData.OrderUID, orderData.TrackNumber, orderData.Entry, orderData.Locale, orderData.InternalSignature,
		orderData.CustomerID, orderData.DeliveryService, orderData.ShardKey, orderData.SmID, orderData.DateCreated, orderData.OofShard,
	)
	n++

	b.Queue(qInsertPayment,
		orderData.OrderUID, orderData.Payment.Transaction, orderData.Payment.RequestID, orderData.Payment.Currency, orderData.Payment.Provider,
		orderData.Payment.Amount, orderData.Payment.PaymentDT, orderData.Payment.Bank, orderData.Payment.DeliveryCost, orderData.Payment.GoodsTotal, orderData.Payment.CustomFee,
	)
	n++

	b.Queue(qInsertDelivery,
		orderData.OrderUID, orderData.Delivery.Name, orderData.Delivery.Phone, orderData.Delivery.Zip, orderData.Delivery.City, orderData.Delivery.Address, orderData.Delivery.Region, orderData.Delivery.Email,
	)
	n++

	for _, it := range orderData.Items {
		b.Queue(qInsertItem,
			orderData.OrderUID, it.ChrtID, it.TrackNumber, it.Price, it.RID, it.Name,
			it.Sale, it.Size, it.TotalPrice, it.NmID, it.Brand, it.Status,
		)
		n++
	}

	return b, n
}
