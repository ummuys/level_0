package repository

import "fmt"

const (
	queryGetN = `
	SELECT
  i.order_uid,
  i.track_number,
  i.entry,
  i.locale,
  i.internal_signature,
  i.customer_id,
  i.delivery_service,
  i.shardkey,
  i.sm_id,
  i.date_created,
  i.oof_shard,

  p.transaction,
  p.request_id,
  p.currency,
  p.provider,
  p.amount,
  p.payment_dt,
  p.bank,
  p.delivery_cost,
  p.goods_total,
  p.custom_fee,

  d.name,
  d.phone,
  d.zip,
  d.city,
  d.address,
  d.region,
  d.email
  
FROM orders.info i
JOIN orders.payments  p ON i.order_uid = p.order_uid
JOIN orders.deliveries d ON i.order_uid = d.order_uid
`

	querySelectItems = `
  SELECT
    chrt_id,
    track_number,
    price,
    rid,
    name,
    sale,
    size,
    total_price,
    nm_id,
    brand,
    status
  FROM orders.items
  WHERE order_uid = $1
  `
)

func createQueryGetN(oUID string, n int) (string, []any) {
	query := queryGetN
	args := []any{}
	i := 1
	if oUID != "" {
		args = append(args, oUID)
		query += fmt.Sprintf(" WHERE i.order_uid = $%d", i)
		i++
	}

	query += fmt.Sprintf(" ORDER BY i.inserted_at DESC")

	if n > 0 {
		args = append(args, n)
		query += fmt.Sprintf(" LIMIT $%d", i)
		i++
	}
	return query, args
}
