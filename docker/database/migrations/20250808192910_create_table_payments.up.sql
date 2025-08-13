CREATE TABLE orders.payments (
    order_uid text primary key references orders.info(order_uid) ON DELETE CASCADE,
    transaction_id TEXT NOT NULL,
    request_id TEXT NULL,
    currency CHAR(3) NOT NULL,
    provider TEXT NOT NULL,
    amount int NOT NULL,
    payment_dt bigint NOT NULL,
    bank TEXT NOT NULL,
    delivery_cost integer NOT NULL,
    goods_total integer NOT NULL,
    custom_fee integer NOT NULL
)