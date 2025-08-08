CREATE TABLE orders.payments (
    order_uid UUID primary key references orders.info(order_uid) ON DELETE CASCADE,
    transaction_id TEXT NOT NULL,
    request_id TEXT NULL,
    currency CHAR(3) NOT NULL,
    provider TEXT NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    payment_dt TIMESTAMPTZ NOT NULL,
    bank TEXT NOT NULL,
    delivery_cost NUMERIC(12,2) NOT NULL,
    goods_total NUMERIC(12,2) NOT NULL,
    custom_fee NUMERIC(12,2) NOT NULL
)