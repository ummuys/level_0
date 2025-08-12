CREATE TABLE orders.items (
  order_uid text NOT NULL REFERENCES orders.info(order_uid) ON DELETE CASCADE,
  chrt_id BIGINT NOT NULL,
  price NUMERIC(12,2) NOT NULL,
  rid TEXT NOT NULL,
  name TEXT NOT NULL,
  sale SMALLINT NOT NULL,
  size TEXT NULL,
  total_price NUMERIC(12,2) NOT NULL,
  nm_id BIGINT NOT NULL,
  brand TEXT NULL,
  status SMALLINT NOT NULL
)