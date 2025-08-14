package repository

import "fmt"

const queryGetN = `
SELECT jsonb_agg(order_json) AS orders
FROM (
  SELECT
    ( to_jsonb(i)
      ||
      jsonb_build_object(
        'payment',  to_jsonb(p) - 'order_uid',
        'delivery', to_jsonb(d) - 'order_uid',
        'items', COALESCE(
          (
            SELECT jsonb_agg(to_jsonb(it) - 'order_uid' ORDER BY it.chrt_id)
            FROM orders.items it
            WHERE it.order_uid = i.order_uid
          ),
          '[]'::jsonb
        )
      )
    ) AS order_json
  FROM orders.info i
  JOIN orders.payments  p on i.order_uid = p.order_uid
  JOIN orders.deliveries d on i.order_uid = d.order_uid
`

func createQueryGetN(oUID string, n int) (string, []any) {
	query := queryGetN
	args := []any{}
	i := 1
	if oUID != "" {
		args = append(args, oUID)
		query += fmt.Sprintf(" WHERE i.order_uid = $%d", i)
		i++
	}
	if n > 1 {
		args = append(args, n)
		query += fmt.Sprintf(" LIMIT $%d", i)
		i++
	}
	query += ` ) query;`
	return query, args
}
