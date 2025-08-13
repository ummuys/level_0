CREATE TABLE orders.items (
  order_uid text NOT NULL REFERENCES orders.info(order_uid) ON DELETE CASCADE,
  chrt_id bigint NOT NULL,
  track_number text NOT NULL,
  price int NOT NULL,
  rid TEXT NOT NULL,
  name TEXT NOT NULL,
  sale int NOT NULL,
  size TEXT NULL,
  total_price int NOT NULL,
  nm_id bigint NOT NULL,
  brand TEXT NULL,
  status SMALLINT NOT NULL
)