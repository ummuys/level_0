CREATE TABLE orders.deliveries (
  order_uid text PRIMARY KEY REFERENCES orders.info(order_uid) ON DELETE CASCADE,
  name TEXT NOT NULL,
  phone TEXT NOT NULL,
  zip TEXT NULL,
  city TEXT NOT NULL,
  address TEXT NOT NULL,
  region TEXT NULL,
  email TEXT NULL
)