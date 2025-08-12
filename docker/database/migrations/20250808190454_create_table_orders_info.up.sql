CREATE TABLE orders.info (
    order_uid text primary key,
    track_number text NOT NULL,
    entry text NOT NULL,
    locale text NOT NULL,
    internal_signature text NOT NULL,
    customer_id text NOT NULL,
    delivery_service text NOT NULL,
    shardkey text NOT NULL,
    sm_id int NOT NULL,
    date_created TIMESTAMPTZ NOT NULL,
    oof_shard int NOT NULL
);