CREATE OR REPLACE FUNCTION orders.insert_order(data jsonb)
RETURNS text
LANGUAGE plpgsql
SECURITY INVOKER
STRICT
AS $$
    DECLARE 
        oUID text;
    BEGIN


    --TODO: ДОДЕЛАТЬ ТАБЛИЦУ
        --main table
        INSERT INTO orders.info (order_uid, track_number, entry, locale, internal_signature,
                                 customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        VALUES (data->>'order_uid', data->>'track_number', data->>'entry', data->>'locale', data->>'internal_signature',
                data->>'customer_id', data->>'delivery_service', data->>'shardkey',
                (data->>'sm_id')::int, (data->>'date_created')::timestamptz, data->>'oof_shard')

        orders.payments (
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